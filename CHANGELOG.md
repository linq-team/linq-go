# Changelog

## [0.34.2](https://github.com/linq-team/linq-go/compare/v0.34.1...v0.34.2) (2026-08-17)


### Bug Fixes

* clarify health check guidance for conversations and volume ([94018df](https://github.com/linq-team/linq-go/commit/94018df8f0d0f4aac7e659b748ab642a1559cfa0))

## [0.34.1](https://github.com/linq-team/linq-go/compare/v0.34.0...v0.34.1) (2026-08-17)


### Bug Fixes

* clarify poll webhook added_options field behavior ([0c9cf58](https://github.com/linq-team/linq-go/commit/0c9cf580a488c613de5614e350e886ad8cf11b10))

## [0.34.0](https://github.com/linq-team/linq-go/compare/v0.33.1...v0.34.0) (2026-08-14)


### Features

* add chat background and poll features ([8537c5e](https://github.com/linq-team/linq-go/commit/8537c5e35ed0b267fb861e5c131b49142dd14e5e))

## [0.33.1](https://github.com/linq-team/linq-go/compare/v0.33.0...v0.33.1) (2026-08-14)


### Bug Fixes

* clarify chat health status guidance and opt-out behavior ([b7bb8f1](https://github.com/linq-team/linq-go/commit/b7bb8f1043d715bdafc637c70b715860aff06607))

## [0.33.0](https://github.com/linq-team/linq-go/compare/v0.32.1...v0.33.0) (2026-08-14)


### Features

* add 403 error responses for ineligible phone numbers ([18968b0](https://github.com/linq-team/linq-go/commit/18968b076e832e276b83f4d2d16f8a5d60b9ab3d))

## [0.32.1](https://github.com/linq-team/linq-go/compare/v0.32.0...v0.32.1) (2026-08-13)


### Bug Fixes

* clarify address format requirements for messaging checks ([f8ba08d](https://github.com/linq-team/linq-go/commit/f8ba08d739abc5a97e595969e133e24c2bfccd7a))

## [0.32.0](https://github.com/linq-team/linq-go/compare/v0.31.1...v0.32.0) (2026-08-13)


### Features

* add detailed 503 errors and response fields for capability checks ([1e1eeef](https://github.com/linq-team/linq-go/commit/1e1eeef2b4ec35bfefcea76b861436f577b06e6c))

## [0.31.1](https://github.com/linq-team/linq-go/compare/v0.31.0...v0.31.1) (2026-08-12)


### Chores

* add auto-merge workflow for release prs ([173ceff](https://github.com/linq-team/linq-go/commit/173ceff52b075889c86c5ea0fd827153343b4b1d))


### Documentation

* add parameter descriptions for phone number audit endpoints ([173ceff](https://github.com/linq-team/linq-go/commit/173ceff52b075889c86c5ea0fd827153343b4b1d))

## [0.31.0](https://github.com/linq-team/linq-go/compare/v0.30.0...v0.31.0) (2026-08-12)


### Features

* add experiences api for native imessage cards ([a451335](https://github.com/linq-team/linq-go/commit/a4513357b7197f22f38d9dcd4c838fd0e1f71e90))
* add line reputation audit endpoints and report types ([a451335](https://github.com/linq-team/linq-go/commit/a4513357b7197f22f38d9dcd4c838fd0e1f71e90))
* rename agentkit to experience in card api ([a451335](https://github.com/linq-team/linq-go/commit/a4513357b7197f22f38d9dcd4c838fd0e1f71e90))


### Bug Fixes

* add example responses and parameter descriptions to experiences endpoints ([a451335](https://github.com/linq-team/linq-go/commit/a4513357b7197f22f38d9dcd4c838fd0e1f71e90))
* clarify api endpoints in documentation examples ([a451335](https://github.com/linq-team/linq-go/commit/a4513357b7197f22f38d9dcd4c838fd0e1f71e90))

## [0.30.0](https://github.com/linq-team/linq-go/compare/v0.29.1...v0.30.0) (2026-08-10)


### Features

* add blocked handles api for managing inbound message filtering ([c554139](https://github.com/linq-team/linq-go/commit/c554139366c7cfd7257a65fe410b54e97aabcf5b))
* add diagnostic fields to message failure webhook events ([c554139](https://github.com/linq-team/linq-go/commit/c554139366c7cfd7257a65fe410b54e97aabcf5b))
* regenerate SDKs from updated API spec ([c554139](https://github.com/linq-team/linq-go/commit/c554139366c7cfd7257a65fe410b54e97aabcf5b))


### Bug Fixes

* clarify opted-out status clearing behavior in chat health ([c554139](https://github.com/linq-team/linq-go/commit/c554139366c7cfd7257a65fe410b54e97aabcf5b))
* clarify stop keyword behavior and matching rules ([c554139](https://github.com/linq-team/linq-go/commit/c554139366c7cfd7257a65fe410b54e97aabcf5b))
* clarify stop keyword behavior for multi-chat replies ([c554139](https://github.com/linq-team/linq-go/commit/c554139366c7cfd7257a65fe410b54e97aabcf5b))
* expand webhook error code documentation for failure events ([c554139](https://github.com/linq-team/linq-go/commit/c554139366c7cfd7257a65fe410b54e97aabcf5b))


### Documentation

* clarify presigned url behavior and configuration ([c554139](https://github.com/linq-team/linq-go/commit/c554139366c7cfd7257a65fe410b54e97aabcf5b))

## [0.29.1](https://github.com/linq-team/linq-go/compare/v0.29.0...v0.29.1) (2026-08-04)


### Bug Fixes

* flatten poll subresources to avoid go compile errors ([ddf0603](https://github.com/linq-team/linq-go/commit/ddf0603d2e5e1a2fa4d1ff62a9b86970fbeeefe3))

## [0.29.0](https://github.com/linq-team/linq-go/compare/v0.28.0...v0.29.0) (2026-08-04)


### Features

* add action field to message content for app experiences ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))
* add agentcard payment provider api ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))
* add exclude_from parameter to control line selection ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))
* add reconciled_at field to message objects ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))
* support url-type fields in message action parameters ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))


### Bug Fixes

* clarify contact card creation and update behavior ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))
* clarify idempotency key behavior with deleted messages ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))
* clarify message.failed webhook delivery behavior ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))
* clarify opt-out keyword matching requirements for chat health ([10be1a9](https://github.com/linq-team/linq-go/commit/10be1a9b39c06e9eba65c49be83c652bd19ab947))

## [0.28.0](https://github.com/linq-team/linq-go/compare/v0.27.0...v0.28.0) (2026-07-26)


### Features

* regenerate SDKs from updated API spec ([#54](https://github.com/linq-team/linq-go/issues/54)) ([665e91b](https://github.com/linq-team/linq-go/commit/665e91b975c68adb761730fb7bd2220880e1cf7a))

## [0.27.0](https://github.com/linq-team/linq-go/compare/v0.26.1...v0.27.0) (2026-07-26)


### Features

* regenerate SDKs from updated API spec ([#52](https://github.com/linq-team/linq-go/issues/52)) ([cd18494](https://github.com/linq-team/linq-go/commit/cd18494310644a5c58683615c69440fc04d9f82e))

## [0.26.1](https://github.com/linq-team/linq-go/compare/v0.26.0...v0.26.1) (2026-07-06)


### Documentation

* clarify phone line reputation status descriptions ([#50](https://github.com/linq-team/linq-go/issues/50)) ([5855d94](https://github.com/linq-team/linq-go/commit/5855d94a8925e32bf4de5abe3dac5e25b0928c3c))

## [0.26.0](https://github.com/linq-team/linq-go/compare/v0.25.0...v0.26.0) (2026-06-21)


### Features

* phone line reputation + group chat icon ([#48](https://github.com/linq-team/linq-go/issues/48)) ([e5339a6](https://github.com/linq-team/linq-go/commit/e5339a6887112b331232492db41d398b0871f411))

## [0.25.0](https://github.com/linq-team/linq-go/compare/v0.24.0...v0.25.0) (2026-06-10)


### Features

* regenerate SDKs from updated API spec ([#45](https://github.com/linq-team/linq-go/issues/45)) ([14e4a51](https://github.com/linq-team/linq-go/commit/14e4a51270ff7f9ebdc48b855354259001f64bfd))


### Bug Fixes

* **ci:** parse release PR number in shell, not fromJSON in env ([#42](https://github.com/linq-team/linq-go/issues/42)) ([20fef3b](https://github.com/linq-team/linq-go/commit/20fef3b2c94c6f9e969469a86a18d5a23afd2710))
