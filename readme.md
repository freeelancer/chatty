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

Some encryption code can be found x/chat/client/cli. Currently this module only supports rsa key pairs pkcs1v15. It can be upgraded to support multiple different keys as ultimately encryption and decryption is handled on frontend applications.

Users generate rsa keys and store their public keys in the blockchain state so other users are able to access it to encrypt messages sent to them. 

Group conversation creators sets the public key for the group conversation on creation (this is optional). The creator will need to find another way to send the participants the private key for this to work effectively.

### commands

```

```
```
chattyd tx chat create-group-conversation <chat_name> <"list of participants separated by space"> <message> <path the pubkey> --from <alice/bob> --keyring-backend test
```

### Future Improvements

- Allow state updates to group conversation like adding/updating of pubkey, participants, admin
- State pruning of messages passed a certain duration. This can be either done through beginblockers or at the beginning of every createMessage transaction


