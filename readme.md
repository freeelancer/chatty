# chatty
**chatty** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Chat Module

The Chat Module is designed such that users can send messages either directly to other users or create group chats where they can speak to multiple participants like any other web chat application. Data is saved in the state, it can also be retrieved from the message data of archival nodes if the state is pruned.

In order to encrypt the messages, the encryption should be done off chain on users own local machines. This is to ensure the privacy of the messages. Although it is possible to encrypt the data with public keys on chain, this is pointless as the call message data are public.

Some encryption code can be found x/chat/client/cli/rsa.go. Currently the cli only supports rsa key pairs pkcs1v15. It can be upgraded to support multiple different keys as ultimately encryption and decryption is handled on frontend applications.

Users generate rsa keys and store their public keys in the blockchain state so other users are able to access it to encrypt messages sent to them. 

Group conversation creators sets the public key for the group conversation on creation (this is optional). The creator will need to find another way to send the participants the private key for this to work effectively.

Cli commands that create messages takes in the <pubkey pem filepath> which can be empty. This would mean the message will not be encrypted.

### Unit Tests
```
go test $(pwd)/x/chat/keeper 
```

### Setup rsa PKCS1 private public key pair

```
ssh-keygen -b 2048 -t rsa -m pem -f ~/.ssh/chatty_rsa
ssh-keygen -f ~/.ssh/chatty_rsa.pub -m 'PEM' -e > ~/.ssh/chatty_rsa.pub.pem
```

### Transaction Commands

```
chattyd tx chat update-pubkey ~/.ssh/chatty_rsa.pub.pem --from alice --keyring-backend test
```
```
chattyd tx chat update-pubkey <pubkey pem filepath> --from <alice/bob> --keyring-backend test
```

Updates the pubkey for user. This will be shown to public.


```
chattyd tx chat create-chat-message cosmos1g7s2hmkun7548huyepkdvy4pw80vjv9v6qn6qt "hi" ~/.ssh/chatty_rsa.pub.pem --from alice --keyring-backend test
```
```
chattyd tx chat create-chat-message <receiver bech32 address> <message> <pubkey pem filepath> --from alice --keyring-backend test
```

Create message from user to receiver, will not be encrypted if filepath is set as "".


```
chattyd tx chat create-group-conversation "nice group" "cosmos1g7s2hmkun7548huyepkdvy4pw80vjv9v6qn6qt" "hi" "" --from alice --keyring-backend test
```
```
chattyd tx chat create-group-conversation <chat_name> <"list of participants separated by space"> <message> <pubkey pem filepath> --from <alice/bob> --keyring-backend test
```

Creates group conversation, message can be "" which will create the group without an initial message

```
chattyd tx chat create-group-conversation-message 1 "hello everyone" "" --from alice --keyring-backend test
```
```
chattyd tx chat create-group-conversation-message <group conversation id> <message> <pubkey pem filepath> --from alice --keyring-backend test
```

Creates message for existing group conversation


### Query Commands
```
chattyd q chat params
```
http://localhost:1317/chatty/chat/params

Get Params which contains they current group_conversation_counter (which is one more than total number of group conversations)

```
chattyd q chat pubkeys
```
http://localhost:1317/chatty/chat/pubkey/

Get all pubkeys

```
chattyd q chat pubkey <address>
```
http://localhost:1317/chatty/chat/pubkey/{address}

Get pubkey of address

```
chattyd q chat conversations <address>
```
http://localhost:1317/chatty/chat/conversation/{address}

Get all one-to-one conversations of address

```
chattyd q chat conversation <addressA> <addressB>
```
http://localhost:1317/chatty/chat/conversation/{address_a}/{address_b}

Get one-to-one conversation between two addresses

```
chattyd q chat group-conversations-by-address <address>
```
http://localhost:1317/chatty/chat/group_conversation/address/{address}

Get all group conversations address is in

```
chattyd q chat group-conversation-by-id <id>
```
http://localhost:1317/chatty/chat/group_conversation/id/{id}

Get group conversation of id


### Future Improvements

- Fix export and init genesis, add test cases
- Allow state updates to group conversation like adding/updating of pubkey, participants, admin
- State pruning of messages passed a certain duration. This can be either done through beginblockers or at the beginning of every createMessage transaction
- Use another way to store one-to-one conversations between two users. Current way stores each conversation twice under two different keys (addressA+addressB and addressB+addressA). It is done this way to faciliate quicker queries of getting all conversations of a single user.


