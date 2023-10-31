package keeper_test

import (
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "chatty/testutil/keeper"
	"chatty/x/chat/keeper"
	"chatty/x/chat/types"

	"chatty/testutil/sample"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, sdk.Context) {
	k, ctx := keepertest.ChatKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

func TestMsgUpdatePubkey(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	alice := sample.AccAddress()

	pemData, err := GetPemDataBs("test_key.pub.pem")
	if err != nil {
		panic(err)
	}

	// default params
	testCases := []struct {
		name      string
		input     *types.MsgUpdatePubkey
		expErr    bool
		expErrMsg string
	}{
		{
			name: "invalid pubkey",
			input: &types.MsgUpdatePubkey{
				Creator: alice,
				Pubkey:  "",
			},
			expErr:    true,
			expErrMsg: "Error decoding PEM block: invalid request",
		},
		{
			name: "valid pubkey",
			input: &types.MsgUpdatePubkey{
				Creator: alice,
				Pubkey:  hex.EncodeToString(pemData),
			},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.UpdatePubkey(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
				pubkey, found := k.GetPubkey(wctx, tc.input.Creator)
				require.True(t, found)
				require.Equal(t, tc.input.Pubkey, pubkey.Key)
			}
		})
	}
}

func GetPemDataBs(filePath string) ([]byte, error) {
	publicKeyFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer publicKeyFile.Close()

	// Read the public key from the PEM file
	pemData, err := io.ReadAll(publicKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("Error decoding public key")
	}
	return pemData, nil
}

func TestCreateChatMessage(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	alice := sample.AccAddress()
	bob := sample.AccAddress()

	// default params
	testCases := []struct {
		name      string
		input     *types.MsgCreateChatMessage
		expErr    bool
		expErrMsg string
	}{
		{
			name: "empty message",
			input: &types.MsgCreateChatMessage{
				Creator:   alice,
				Receiver:  bob,
				Message:   "",
				Encrypted: false,
			},
			expErr:    true,
			expErrMsg: "message is empty: invalid request",
		},
		{
			name: "same address",
			input: &types.MsgCreateChatMessage{
				Creator:   alice,
				Receiver:  alice,
				Message:   "hi",
				Encrypted: false,
			},
			expErr:    true,
			expErrMsg: "creator and receiver are the same: invalid request",
		},
		{
			name: "invalid address",
			input: &types.MsgCreateChatMessage{
				Creator:   alice,
				Receiver:  "aaaa",
				Message:   "hi",
				Encrypted: false,
			},
			expErr:    true,
			expErrMsg: "invalid receiver address (decoding bech32 failed: invalid bech32 string length 4): invalid address",
		},
		{
			name: "valid",
			input: &types.MsgCreateChatMessage{
				Creator:   alice,
				Receiver:  bob,
				Message:   "hi",
				Encrypted: false,
			},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.CreateChatMessage(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
				conversation, found := k.GetConversation(wctx, tc.input.Creator, tc.input.Receiver)
				require.True(t, found)
				chatMessage := conversation.Messages[len(conversation.Messages)-1]
				require.Equal(t, tc.input.Message, chatMessage.Message)
				require.Equal(t, tc.input.Encrypted, chatMessage.Encrypted)
				require.Equal(t, tc.input.Creator, chatMessage.Sender)
				require.Contains(t, []string{conversation.ChatterA, conversation.ChatterB}, tc.input.Receiver)
			}
		})
	}
}

func TestCreateGroupConversation(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	alice := sample.AccAddress()
	bob := sample.AccAddress()
	charlie := sample.AccAddress()

	pemData, err := GetPemDataBs("test_key.pub.pem")
	if err != nil {
		panic(err)
	}
	pemDataStr := hex.EncodeToString(pemData)

	// default params
	testCases := []struct {
		name      string
		input     *types.MsgCreateGroupConversation
		expErr    bool
		expErrMsg string
	}{
		{
			name: "empty name",
			input: &types.MsgCreateGroupConversation{
				Creator:      alice,
				Name:         "",
				Participants: []string{bob, charlie},
				Pubkey:       "",
				Message:      "",
				Encrypted:    false,
			},
			expErr:    true,
			expErrMsg: "name cannot be empty: invalid request",
		},
		{
			name: "same address",
			input: &types.MsgCreateGroupConversation{
				Creator:      alice,
				Name:         "Test",
				Participants: []string{alice},
				Pubkey:       "",
				Message:      "",
				Encrypted:    false,
			},
			expErr:    true,
			expErrMsg: "participants only has creator: invalid request",
		},
		{
			name: "invalid address",
			input: &types.MsgCreateGroupConversation{
				Creator:      alice,
				Name:         "Test",
				Participants: []string{bob, "cccc"},
				Pubkey:       "",
				Message:      "",
				Encrypted:    false,
			},
			expErr:    true,
			expErrMsg: "invalid participant address (decoding bech32 failed: invalid bech32 string length 4): invalid address",
		},
		{
			name: "valid",
			input: &types.MsgCreateGroupConversation{
				Creator:      alice,
				Name:         "Test",
				Participants: []string{bob, charlie},
				Pubkey:       "",
				Message:      "",
				Encrypted:    false,
			},
			expErr:    false,
			expErrMsg: "",
		},
		{
			name: "valid pubkey",
			input: &types.MsgCreateGroupConversation{
				Creator:      alice,
				Name:         "Test",
				Participants: []string{bob, charlie},
				Pubkey:       pemDataStr,
				Message:      "",
				Encrypted:    false,
			},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.CreateGroupConversation(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
				id := k.GetParams(wctx).GroupConversationCounter - int64(1)
				conversation, found := k.GetGroupConversation(wctx, id)
				require.True(t, found)
				if len(conversation.Messages) > 1 {
					chatMessage := conversation.Messages[len(conversation.Messages)-1]
					require.Equal(t, tc.input.Message, chatMessage.Message)
					require.Equal(t, tc.input.Encrypted, chatMessage.Encrypted)
					require.Equal(t, tc.input.Creator, chatMessage.Sender)
				}
				if tc.input.Pubkey == "" {
					require.Nil(t, conversation.PubKey)
				} else {
					require.Equal(t, tc.input.Pubkey, conversation.PubKey.Key)
					require.Equal(t, tc.input.Creator, conversation.PubKey.Address)
				}

				require.ElementsMatch(t, []string{alice, bob, charlie}, conversation.Participants)
				require.Equal(t, tc.input.Creator, conversation.Admin)
				require.Equal(t, id, conversation.Id)
				require.Equal(t, tc.input.Name, conversation.Name)
			}
		})
	}
}

func TestCreateGroupConversationMessage(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	alice := sample.AccAddress()
	bob := sample.AccAddress()
	charlie := sample.AccAddress()

	pemData, err := GetPemDataBs("test_key.pub.pem")
	if err != nil {
		panic(err)
	}
	pemDataStr := hex.EncodeToString(pemData)

	_, err = ms.CreateGroupConversation(wctx, &types.MsgCreateGroupConversation{
		Creator:      alice,
		Name:         "Test",
		Participants: []string{bob, charlie},
		Pubkey:       pemDataStr,
		Message:      "",
		Encrypted:    false,
	})
	require.NoError(t, err)

	_, err = ms.CreateGroupConversation(wctx, &types.MsgCreateGroupConversation{
		Creator:      alice,
		Name:         "Test2",
		Participants: []string{charlie},
		Pubkey:       pemDataStr,
		Message:      "",
		Encrypted:    false,
	})
	require.NoError(t, err)

	// default params
	testCases := []struct {
		name      string
		input     *types.MsgCreateGroupConversationMessage
		expErr    bool
		expErrMsg string
	}{
		{
			name: "empty message",
			input: &types.MsgCreateGroupConversationMessage{
				Creator:        bob,
				ConversationId: 1,
				Message:        "",
				Encrypted:      false,
			},
			expErr:    true,
			expErrMsg: "message is empty: invalid request",
		},
		{
			name: "not found id",
			input: &types.MsgCreateGroupConversationMessage{
				Creator:        bob,
				ConversationId: 5,
				Message:        "hello",
				Encrypted:      false,
			},
			expErr:    true,
			expErrMsg: "group conversation with id 5 does not exist",
		},
		{
			name: "address not in participants",
			input: &types.MsgCreateGroupConversationMessage{
				Creator:        bob,
				ConversationId: 2,
				Message:        "hello",
				Encrypted:      false,
			},
			expErr:    true,
			expErrMsg: "is not a participant of group conversation with id 2",
		},
		{
			name: "valid",
			input: &types.MsgCreateGroupConversationMessage{
				Creator:        bob,
				ConversationId: 1,
				Message:        "hello",
				Encrypted:      false,
			},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.CreateGroupConversationMessage(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
				conversation, found := k.GetGroupConversation(wctx, 1)
				require.True(t, found)

				chatMessage := conversation.Messages[len(conversation.Messages)-1]
				require.Equal(t, tc.input.Message, chatMessage.Message)
				require.Equal(t, tc.input.Encrypted, chatMessage.Encrypted)
				require.Equal(t, tc.input.Creator, chatMessage.Sender)
				require.Equal(t, tc.input.ConversationId, conversation.Id)
			}
		})
	}
}
