# Shared Params Types

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#ReactionType">ReactionType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#ServiceType">ServiceType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#TextDecorationParam">TextDecorationParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#ChatHandle">ChatHandle</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#LinkPartResponse">LinkPartResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#MediaPartResponse">MediaPartResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#ReactionType">ReactionType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#ServiceType">ServiceType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#TextDecoration">TextDecoration</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/shared#TextPartResponse">TextPartResponse</a>

# Chats

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#LinkPartParam">LinkPartParam</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MediaPartParam">MediaPartParam</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageContentParam">MessageContentParam</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#TextPartParam">TextPartParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Chat">Chat</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatNewResponse">ChatNewResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatUpdateResponse">ChatUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatLeaveChatResponse">ChatLeaveChatResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatSendVoicememoResponse">ChatSendVoicememoResponse</a>

Methods:

- <code title="post /v3/chats">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatNewParams">ChatNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatNewResponse">ChatNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/chats/{chatId}">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/chats/{chatId}">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatUpdateParams">ChatUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatUpdateResponse">ChatUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/leave">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.LeaveChat">LeaveChat</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatLeaveChatResponse">ChatLeaveChatResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/chats">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.ListChats">ListChats</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatListChatsParams">ChatListChatsParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/packages/pagination#ListChatsPagination">ListChatsPagination</a>[<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Chat">Chat</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/read">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.MarkAsRead">MarkAsRead</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/chats/{chatId}/voicememo">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.SendVoicememo">SendVoicememo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatSendVoicememoParams">ChatSendVoicememoParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatSendVoicememoResponse">ChatSendVoicememoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/share_contact_card">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.ShareContactCard">ShareContactCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Participants

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantAddResponse">ChatParticipantAddResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantRemoveResponse">ChatParticipantRemoveResponse</a>

Methods:

- <code title="post /v3/chats/{chatId}/participants">client.Chats.Participants.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantAddParams">ChatParticipantAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantAddResponse">ChatParticipantAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/chats/{chatId}/participants">client.Chats.Participants.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantRemoveParams">ChatParticipantRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantRemoveResponse">ChatParticipantRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Typing

Methods:

- <code title="post /v3/chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatTypingService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v3/chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatTypingService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SentMessage">SentMessage</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageSendResponse">ChatMessageSendResponse</a>

Methods:

- <code title="get /v3/chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageListParams">ChatMessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/packages/pagination#ListMessagesPagination">ListMessagesPagination</a>[<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageSendParams">ChatMessageSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageSendResponse">ChatMessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageEffectParam">MessageEffectParam</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReplyToParam">ReplyToParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageEffect">MessageEffect</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReplyTo">ReplyTo</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageAddReactionResponse">MessageAddReactionResponse</a>

Methods:

- <code title="get /v3/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageUpdateParams">MessageUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/messages/{messageId}/reactions">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.AddReaction">AddReaction</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageAddReactionParams">MessageAddReactionParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageAddReactionResponse">MessageAddReactionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/messages/{messageId}/thread">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.ListMessagesThread">ListMessagesThread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageListMessagesThreadParams">MessageListMessagesThreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go/packages/pagination#ListMessagesPagination">ListMessagesPagination</a>[<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Attachments

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SupportedContentType">SupportedContentType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SupportedContentType">SupportedContentType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentNewResponse">AttachmentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentGetResponse">AttachmentGetResponse</a>

Methods:

- <code title="post /v3/attachments">client.Attachments.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentNewParams">AttachmentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentNewResponse">AttachmentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/attachments/{attachmentId}">client.Attachments.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attachmentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentGetResponse">AttachmentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Phonenumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhonenumberListResponse">PhonenumberListResponse</a>

Methods:

- <code title="get /v3/phonenumbers">client.Phonenumbers.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhonenumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhonenumberListResponse">PhonenumberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PhoneNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhoneNumberListResponse">PhoneNumberListResponse</a>

Methods:

- <code title="get /v3/phone_numbers">client.PhoneNumbers.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhoneNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhoneNumberListResponse">PhoneNumberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WebhookEvents

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventType">WebhookEventType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventType">WebhookEventType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventListResponse">WebhookEventListResponse</a>

Methods:

- <code title="get /v3/webhook-events">client.WebhookEvents.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventListResponse">WebhookEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WebhookSubscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscription">WebhookSubscription</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionNewResponse">WebhookSubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionListResponse">WebhookSubscriptionListResponse</a>

Methods:

- <code title="post /v3/webhook-subscriptions">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionNewParams">WebhookSubscriptionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionNewResponse">WebhookSubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionUpdateParams">WebhookSubscriptionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/webhook-subscriptions">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionListResponse">WebhookSubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Capability

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#HandleCheckParam">HandleCheckParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#HandleCheckResponse">HandleCheckResponse</a>

Methods:

- <code title="post /v3/capability/check_imessage">client.Capability.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityService.CheckiMessage">CheckiMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityCheckiMessageParams">CapabilityCheckiMessageParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#HandleCheckResponse">HandleCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/capability/check_rcs">client.Capability.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityService.CheckRCS">CheckRCS</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityCheckRCSParams">CapabilityCheckRCSParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#HandleCheckResponse">HandleCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageEventV2">MessageEventV2</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReactionEventBase">ReactionEventBase</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SchemasMediaPartResponse">SchemasMediaPartResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SchemasMessageEffect">SchemasMessageEffect</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SchemasTextPartResponse">SchemasTextPartResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageSentWebhookEvent">MessageSentWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageReceivedWebhookEvent">MessageReceivedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageReadWebhookEvent">MessageReadWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageDeliveredWebhookEvent">MessageDeliveredWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageFailedWebhookEvent">MessageFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageEditedWebhookEvent">MessageEditedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReactionAddedWebhookEvent">ReactionAddedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReactionRemovedWebhookEvent">ReactionRemovedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ParticipantAddedWebhookEvent">ParticipantAddedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ParticipantRemovedWebhookEvent">ParticipantRemovedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatCreatedWebhookEvent">ChatCreatedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatGroupNameUpdatedWebhookEvent">ChatGroupNameUpdatedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatGroupIconUpdatedWebhookEvent">ChatGroupIconUpdatedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatGroupNameUpdateFailedWebhookEvent">ChatGroupNameUpdateFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatGroupIconUpdateFailedWebhookEvent">ChatGroupIconUpdateFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatTypingIndicatorStartedWebhookEvent">ChatTypingIndicatorStartedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatTypingIndicatorStoppedWebhookEvent">ChatTypingIndicatorStoppedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhoneNumberStatusUpdatedWebhookEvent">PhoneNumberStatusUpdatedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#EventsWebhookEventUnion">EventsWebhookEventUnion</a>

# ContactCard

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SetContactCard">SetContactCard</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ContactCardGetResponse">ContactCardGetResponse</a>

Methods:

- <code title="post /v3/contact_card">client.ContactCard.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ContactCardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ContactCardNewParams">ContactCardNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SetContactCard">SetContactCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/contact_card">client.ContactCard.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ContactCardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ContactCardGetParams">ContactCardGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ContactCardGetResponse">ContactCardGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/contact_card">client.ContactCard.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ContactCardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ContactCardUpdateParams">ContactCardUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqgo</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SetContactCard">SetContactCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
