[![Go Reference](https://pkg.go.dev/badge/github.com/thefuga/hexa-mailbox.svg)](https://pkg.go.dev/github.com/thefuga/hexa-mailbox)
------
# Hexa Mailbox
Simple conceptual project to apply the mailbox patter on an ports/adapters architecture.

The ideias here are - probably - not ideal, but the challange (and the point of this repo) is to communicate between ports and adapters with the same structural type asynchronously without serializations while avoidingcircular imports.

To achieve this, all comunication is done via mailboxes, represented by channels.
When synchronouys communication is needed, two mailboxes must be used: one for sending messages and other to receive responses.
The dependencies are resolved on main.

All names used here are merely conceptual.

Implementation details can be found as godocs on each package.
