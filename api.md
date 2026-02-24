# Chats

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageContentParam">MessageContentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Chat">Chat</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatNewResponse">ChatNewResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatListResponse">ChatListResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatSendVoicememoResponse">ChatSendVoicememoResponse</a>

Methods:

- <code title="post /v3/chats">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatNewParams">ChatNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatNewResponse">ChatNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/chats/{chatId}">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/chats/{chatId}">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatUpdateParams">ChatUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/chats">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatListParams">ChatListParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatListResponse">ChatListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/read">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.MarkAsRead">MarkAsRead</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/chats/{chatId}/voicememo">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.SendVoicememo">SendVoicememo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatSendVoicememoParams">ChatSendVoicememoParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatSendVoicememoResponse">ChatSendVoicememoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/share_contact_card">client.Chats.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatService.ShareContactCard">ShareContactCard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Participants

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantAddResponse">ChatParticipantAddResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantRemoveResponse">ChatParticipantRemoveResponse</a>

Methods:

- <code title="post /v3/chats/{chatId}/participants">client.Chats.Participants.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantAddParams">ChatParticipantAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantAddResponse">ChatParticipantAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/chats/{chatId}/participants">client.Chats.Participants.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantRemoveParams">ChatParticipantRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatParticipantRemoveResponse">ChatParticipantRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Typing

Methods:

- <code title="post /v3/chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatTypingService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v3/chats/{chatId}/typing">client.Chats.Typing.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatTypingService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SentMessage">SentMessage</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageListResponse">ChatMessageListResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageSendResponse">ChatMessageSendResponse</a>

Methods:

- <code title="get /v3/chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageListParams">ChatMessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageListResponse">ChatMessageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/chats/{chatId}/messages">client.Chats.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageSendParams">ChatMessageSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatMessageSendResponse">ChatMessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageEffectParam">MessageEffectParam</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReactionType">ReactionType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReplyToParam">ReplyToParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ChatHandle">ChatHandle</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MediaPart">MediaPart</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageEffect">MessageEffect</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReactionType">ReactionType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#ReplyTo">ReplyTo</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#TextPart">TextPart</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageGetThreadResponse">MessageGetThreadResponse</a>

Methods:

- <code title="get /v3/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/messages/{messageId}">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageDeleteParams">MessageDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/messages/{messageId}/reactions">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.AddReaction">AddReaction</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageAddReactionParams">MessageAddReactionParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#Reaction">Reaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/messages/{messageId}/thread">client.Messages.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageService.GetThread">GetThread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageGetThreadParams">MessageGetThreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#MessageGetThreadResponse">MessageGetThreadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Attachments

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SupportedContentType">SupportedContentType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#SupportedContentType">SupportedContentType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentNewResponse">AttachmentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentGetResponse">AttachmentGetResponse</a>

Methods:

- <code title="post /v3/attachments">client.Attachments.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentNewParams">AttachmentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentNewResponse">AttachmentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/attachments/{attachmentId}">client.Attachments.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attachmentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#AttachmentGetResponse">AttachmentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Phonenumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhonenumberListResponse">PhonenumberListResponse</a>

Methods:

- <code title="get /v3/phonenumbers">client.Phonenumbers.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhonenumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhonenumberListResponse">PhonenumberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PhoneNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhoneNumberListResponse">PhoneNumberListResponse</a>

Methods:

- <code title="get /v3/phone_numbers">client.PhoneNumbers.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhoneNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#PhoneNumberListResponse">PhoneNumberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WebhookEvents

Params Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventType">WebhookEventType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventType">WebhookEventType</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventListResponse">WebhookEventListResponse</a>

Methods:

- <code title="get /v3/webhook-events">client.WebhookEvents.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookEventListResponse">WebhookEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WebhookSubscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscription">WebhookSubscription</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionNewResponse">WebhookSubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionListResponse">WebhookSubscriptionListResponse</a>

Methods:

- <code title="post /v3/webhook-subscriptions">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionNewParams">WebhookSubscriptionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionNewResponse">WebhookSubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionUpdateParams">WebhookSubscriptionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/webhook-subscriptions">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionListResponse">WebhookSubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/webhook-subscriptions/{subscriptionId}">client.WebhookSubscriptions.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#WebhookSubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Capability

Response Types:

- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityCheckImessageResponse">CapabilityCheckImessageResponse</a>
- <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityCheckRcsResponse">CapabilityCheckRcsResponse</a>

Methods:

- <code title="post /v3/capability/check_imessage">client.Capability.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityService.CheckImessage">CheckImessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityCheckImessageParams">CapabilityCheckImessageParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityCheckImessageResponse">CapabilityCheckImessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/capability/check_rcs">client.Capability.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityService.CheckRcs">CheckRcs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityCheckRcsParams">CapabilityCheckRcsParams</a>) (\*<a href="https://pkg.go.dev/github.com/linq-team/linq-go">linqapiv3</a>.<a href="https://pkg.go.dev/github.com/linq-team/linq-go#CapabilityCheckRcsResponse">CapabilityCheckRcsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
