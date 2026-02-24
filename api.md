# Chats

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageContentParam">MessageContentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#Chat">Chat</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatNewResponse">ChatNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatListResponse">ChatListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatSendVoicememoResponse">ChatSendVoicememoResponse</a>

Methods:

- <code title="post /v3/chats">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatNewParams">ChatNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatNewResponse">ChatNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/chats/{chatId}">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/chats/{chatId}">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatUpdateParams">ChatUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/chats">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatListParams">ChatListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatListResponse">ChatListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/read">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatService.MarkAsRead">MarkAsRead</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/chats/{chatId}/voicememo">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatService.SendVoicememo">SendVoicememo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatSendVoicememoParams">ChatSendVoicememoParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatSendVoicememoResponse">ChatSendVoicememoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/share_contact_card">client.Chats.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatService.ShareContactCard">ShareContactCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Participants

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatParticipantAddResponse">ChatParticipantAddResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatParticipantRemoveResponse">ChatParticipantRemoveResponse</a>

Methods:

- <code title="post /v3/chats/{chatId}/participants">client.Chats.Participants.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatParticipantService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatParticipantAddParams">ChatParticipantAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatParticipantAddResponse">ChatParticipantAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/chats/{chatId}/participants">client.Chats.Participants.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatParticipantService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatParticipantRemoveParams">ChatParticipantRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatParticipantRemoveResponse">ChatParticipantRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Typing

Methods:

- <code title="post /v3/chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatTypingService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v3/chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatTypingService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#SentMessage">SentMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatMessageListResponse">ChatMessageListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatMessageSendResponse">ChatMessageSendResponse</a>

Methods:

- <code title="get /v3/chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatMessageListParams">ChatMessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatMessageListResponse">ChatMessageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatMessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatMessageSendParams">ChatMessageSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatMessageSendResponse">ChatMessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageEffectParam">MessageEffectParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ReactionType">ReactionType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ReplyToParam">ReplyToParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ChatHandle">ChatHandle</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MediaPart">MediaPart</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageEffect">MessageEffect</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ReactionType">ReactionType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#ReplyTo">ReplyTo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#TextPart">TextPart</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageGetThreadResponse">MessageGetThreadResponse</a>

Methods:

- <code title="get /v3/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageDeleteParams">MessageDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/messages/{messageId}/reactions">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageService.AddReaction">AddReaction</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageAddReactionParams">MessageAddReactionParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#Reaction">Reaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/messages/{messageId}/thread">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageService.GetThread">GetThread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageGetThreadParams">MessageGetThreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#MessageGetThreadResponse">MessageGetThreadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Attachments

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#SupportedContentType">SupportedContentType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#SupportedContentType">SupportedContentType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#AttachmentNewResponse">AttachmentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#AttachmentGetResponse">AttachmentGetResponse</a>

Methods:

- <code title="post /v3/attachments">client.Attachments.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#AttachmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#AttachmentNewParams">AttachmentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#AttachmentNewResponse">AttachmentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/attachments/{attachmentId}">client.Attachments.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#AttachmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attachmentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#AttachmentGetResponse">AttachmentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Phonenumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#PhonenumberListResponse">PhonenumberListResponse</a>

Methods:

- <code title="get /v3/phonenumbers">client.Phonenumbers.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#PhonenumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#PhonenumberListResponse">PhonenumberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PhoneNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#PhoneNumberListResponse">PhoneNumberListResponse</a>

Methods:

- <code title="get /v3/phone_numbers">client.PhoneNumbers.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#PhoneNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#PhoneNumberListResponse">PhoneNumberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WebhookEvents

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookEventType">WebhookEventType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookEventType">WebhookEventType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookEventListResponse">WebhookEventListResponse</a>

Methods:

- <code title="get /v3/webhook-events">client.WebhookEvents.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookEventListResponse">WebhookEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WebhookSubscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscription">WebhookSubscription</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionNewResponse">WebhookSubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionListResponse">WebhookSubscriptionListResponse</a>

Methods:

- <code title="post /v3/webhook-subscriptions">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionNewParams">WebhookSubscriptionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionNewResponse">WebhookSubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionUpdateParams">WebhookSubscriptionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/webhook-subscriptions">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionListResponse">WebhookSubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#WebhookSubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Capability

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#CapabilityCheckImessageResponse">CapabilityCheckImessageResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#CapabilityCheckRcsResponse">CapabilityCheckRcsResponse</a>

Methods:

- <code title="post /v3/capability/check_imessage">client.Capability.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#CapabilityService.CheckImessage">CheckImessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#CapabilityCheckImessageParams">CapabilityCheckImessageParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#CapabilityCheckImessageResponse">CapabilityCheckImessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/capability/check_rcs">client.Capability.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#CapabilityService.CheckRcs">CheckRcs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#CapabilityCheckRcsParams">CapabilityCheckRcsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/linq-api-v3-go#CapabilityCheckRcsResponse">CapabilityCheckRcsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
