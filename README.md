# Custom Protocol Built on TCP/IP

Welcome to the repository for my custom protocol, built on top of TCP/IP. This protocol provides a simple command-based system for real-time communication between clients and servers. It supports basic operations like registering users, joining channels, messaging, and listing users or channels.

## Features
- **Client Registration**: Allows users to register as clients on the server.
- **Channel Management**: Clients can join or leave communication channels.
- **Messaging**: Enables sending and receiving messages between users and channels.
- **User and Channel Listing**: Clients can query the server for active users and channels.
- **Server Feedback**: Acknowledgement and error handling for all commands.

---

## Commands Overview

| Command | Issued By | Description                                    |
|---------|-----------|------------------------------------------------|
| `REG`   | Client    | Register as a client.                         |
| `JOIN`  | Client    | Join a specific channel.                      |
| `LEAVE` | Client    | Leave the current channel.                    |
| `CHNS`  | Client    | List all available channels.                  |
| `USRS`  | Client    | List all users connected to the server.       |
| `OK`    | Server    | Acknowledge the successful execution of a command. |
| `ERR`   | Server    | Indicate an error in the received command.    |
| `MSG`   | Both      | Send or receive a message to/from a user or channel. |

---

## How It Works

1. **Client Registration (`REG`)**  
   The client starts by sending a `REG` command to the server to register itself.
    ```plaintext
   REG <handle>
    ```
   Example:  
   ```plaintext
   REG @john
2. **Joining a Channel (`JOIN`)**  
    A registered client can join a communication channel using the `JOIN` command.  
      ```plaintext
   JOIN <channel-id>
    ```
   Example:  
   ```plaintext
   JOIN #general
3. **Leaving a Channel (`LEAVE`)**  
    Clients can leave the channel they are currently in using `LEAVE`.  
    ```plaintext
   LEAVE <channel-id>
    ```
   Example:  
   ```plaintext
   LEAVE #general
4. **Messaging (`MSG`)**  
    Messages can be sent to a specific user or channel. Messages from the server also arrive using this command.  
      ```plaintext
   MSG <entity-id> <length>\r\n[payload]
    ```
   Example:  
   ```plaintext
   MSG #general 16\r\nHello everyone!
5. **Listing Channels (`CHNS`)**  
    To get a list of all active channels, clients use the `CHNS` command.  
    **Example**:  
    ```plaintext
    CHNS
    ```

6. **Listing Users (`USRS`)**  
    To see a list of connected users, clients issue the `USRS` command.  
    **Example**:  
    ```plaintext
    USRS
    ```

8. **Acknowledgement and Errors (`OK` and `ERR`)**  
    The server responds with `OK` for successful operations or `ERR` for errors.  
    **Example**:  
    ```plaintext
    ERR Channel ID must begin with #
    ERR Username must begin with @
    ```
