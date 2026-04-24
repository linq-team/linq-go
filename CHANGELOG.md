# Changelog

## 0.14.0 (2026-04-23)

Full Changelog: [v0.13.3...v0.14.0](https://github.com/linq-team/linq-go/compare/v0.13.3...v0.14.0)

### Features

* **api:** expose health_score on chats (BETA) ([fa0433a](https://github.com/linq-team/linq-go/commit/fa0433ae8960429da3a85a9ec27be5977d02d72c))


### Chores

* **internal:** more robust bootstrap script ([5155430](https://github.com/linq-team/linq-go/commit/51554303e3275855eb7be943e126783756d2bd3f))


### Documentation

* **api:** document edit message limits (BUG-7607) ([cf5add2](https://github.com/linq-team/linq-go/commit/cf5add22ef2bcb8dfbcc685cb710550f2539dd65))

## 0.13.3 (2026-04-14)

Full Changelog: [v0.13.2...v0.13.3](https://github.com/linq-team/linq-go/compare/v0.13.2...v0.13.3)

### Documentation

* **openapi:** document typing indicator behavior and limitations ([5e9887a](https://github.com/linq-team/linq-go/commit/5e9887a6c49a94e83355fea1a2ff2efea938dd21))

## 0.13.2 (2026-04-08)

Full Changelog: [v0.13.1...v0.13.2](https://github.com/linq-team/linq-go/compare/v0.13.1...v0.13.2)

### Bug Fixes

* **api-service:** add created_at and make sent_at nullable in SentMessage ([0a4e049](https://github.com/linq-team/linq-go/commit/0a4e0497a1a696ffbdad10d7e7667021cd11f991))
* block SMS group participant changes and fix e2e test failures ([d204608](https://github.com/linq-team/linq-go/commit/d204608f61abc318d7f37e0825295afbb1c7d32c))

## 0.13.1 (2026-04-07)

Full Changelog: [v0.13.0...v0.13.1](https://github.com/linq-team/linq-go/compare/v0.13.0...v0.13.1)

### Bug Fixes

* add SVG support to synapse attachments ([5cc4ad1](https://github.com/linq-team/linq-go/commit/5cc4ad1cd43e789e82bbd48a428061743dae0061))

## 0.13.0 (2026-04-04)

Full Changelog: [v0.12.1...v0.13.0](https://github.com/linq-team/linq-go/compare/v0.12.1...v0.13.0)

### Features

* **api:** config cleanup ([3cb9687](https://github.com/linq-team/linq-go/commit/3cb96871f18e82a6a1b94510b26c6df7f89e3ac1))

## 0.12.1 (2026-04-01)

Full Changelog: [v0.12.0...v0.12.1](https://github.com/linq-team/linq-go/compare/v0.12.0...v0.12.1)

### Documentation

* update contact card API docs with setup and sharing guidance ([69a6fcc](https://github.com/linq-team/linq-go/commit/69a6fccbc328f51bf90f30f4791afbb6e0fd6fd7))

## 0.12.0 (2026-04-01)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/linq-team/linq-go/compare/v0.11.0...v0.12.0)

### Features

* **api:** fix webhook model ([9df7242](https://github.com/linq-team/linq-go/commit/9df724230dac439ec6f7cf84399acc0082a4adf3))
* **internal:** support comma format in multipart form encoding ([ec8156b](https://github.com/linq-team/linq-go/commit/ec8156b30819ad897be0526fcdd869e31e12ed3e))
* PLT(Synapse): Add attachment_id support and tests to voice memo endpoint ([94e805f](https://github.com/linq-team/linq-go/commit/94e805f6d35b6b15a433d2a36443ee1056790bc6))
* Return 403 for group chat typing indicators ([835f8d9](https://github.com/linq-team/linq-go/commit/835f8d9926e4ba5a8a24190ec12f319c3e9e0520))


### Bug Fixes

* fix issue with unmarshaling in some cases ([6e746ef](https://github.com/linq-team/linq-go/commit/6e746ef4ecad30c20d5ce53507d2a16fbda759d6))
* prevent duplicate ? in query params ([38e4bf9](https://github.com/linq-team/linq-go/commit/38e4bf99ec4bf5e2203570e0c9ef036b7262a23a))


### Chores

* **ci:** skip lint on metadata-only changes ([2ef87de](https://github.com/linq-team/linq-go/commit/2ef87def714fdd0d359d6429fb01957c1cac09ff))
* **ci:** support opting out of skipping builds on metadata-only commits ([83b9552](https://github.com/linq-team/linq-go/commit/83b95525205f04bc0304d930f616cc2ca485c85b))
* **client:** fix multipart serialisation of Default() fields ([5db0ada](https://github.com/linq-team/linq-go/commit/5db0ada36b91f99391777fe980f36fde73956d68))
* **internal:** support default value struct tag ([eb74b74](https://github.com/linq-team/linq-go/commit/eb74b743ce88821287533390d4ac4a649c97ca4d))
* remove unnecessary error check for url parsing ([320d34e](https://github.com/linq-team/linq-go/commit/320d34ee77821e77695ef266881a4d578c84798b))
* update docs for api:"required" ([576fb67](https://github.com/linq-team/linq-go/commit/576fb67c7671c53f0006d3d5803c668dd888de6a))

## 0.11.0 (2026-03-24)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/linq-team/linq-go/compare/v0.10.0...v0.11.0)

### Features

* Clean up mess in docs ([72fd9d2](https://github.com/linq-team/linq-go/commit/72fd9d29512b7b1118be62871c4b5183fb5085a4))

## 0.10.0 (2026-03-24)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/linq-team/linq-go/compare/v0.9.0...v0.10.0)

### Features

* BUG: soft delete messages in V3 ([d392bd9](https://github.com/linq-team/linq-go/commit/d392bd91c08932a74669796512a510d9e4a91253))
* Pdev 6191 facetime orchestrator hub api service call service webhook ([f97d6e1](https://github.com/linq-team/linq-go/commit/f97d6e15304545a0390c7e7139cade363ce376c2))


### Bug Fixes

* **webhook:** add NATS BackOff retry + tighten squawk linter ([02dd2d6](https://github.com/linq-team/linq-go/commit/02dd2d6df8921a9756195cf9df6b33f445df0c0c))


### Chores

* **internal:** update gitignore ([a468c50](https://github.com/linq-team/linq-go/commit/a468c5086b98ac16e90c5b8ebdb5e773aed81a94))

## 0.9.0 (2026-03-20)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/linq-team/linq-go/compare/v0.8.0...v0.9.0)

### Features

* add per-line phone number filtering for webhook subscriptions ([151d41a](https://github.com/linq-team/linq-go/commit/151d41ac9c43a5349273e78594f340649eb2c71e))


### Bug Fixes

* return link part type in API responses and webhooks ([8674987](https://github.com/linq-team/linq-go/commit/86749878bb3da150f361a6407d2d20ee1f629998))

## 0.8.0 (2026-03-19)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/linq-team/linq-go/compare/v0.7.0...v0.8.0)

### Features

* **api:** manual updates ([f78d201](https://github.com/linq-team/linq-go/commit/f78d201fb059e0ee4ce1a3ce997e5f973f990699))
* **api:** update config ([72e87bd](https://github.com/linq-team/linq-go/commit/72e87bdad5592a359d5beea55c49e054724db0ae))

## 0.7.0 (2026-03-19)

Full Changelog: [v0.6.1...v0.7.0](https://github.com/linq-team/linq-go/compare/v0.6.1...v0.7.0)

### Features

* BUG: support rich media ddscan in links ([c49b87f](https://github.com/linq-team/linq-go/commit/c49b87f3a6b255989af02c948dd2dacb5f6d8cf4))
* PLT(Synapse): Add Content-Type validation for outbound presigned URL uploads ([33ea384](https://github.com/linq-team/linq-go/commit/33ea384df63694bcc96016ba02b72870f9c51b1e))

## 0.6.1 (2026-03-17)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/linq-team/linq-go/compare/v0.6.0...v0.6.1)

### Bug Fixes

* enforce server-side authorization on DELETE /v3/messages/{messageId} ([b60931f](https://github.com/linq-team/linq-go/commit/b60931f716ef325d3bd97ed03706b304834815a5))
* **openapi:** correct schema errors and example inconsistencies ([82dc722](https://github.com/linq-team/linq-go/commit/82dc722bfee331da0689603812d221a7cb56e37e))


### Chores

* **internal:** tweak CI branches ([95c472f](https://github.com/linq-team/linq-go/commit/95c472f4d80bee944d706d3d7846a8114805ffd0))

## 0.6.0 (2026-03-12)

Full Changelog: [v0.5.1...v0.6.0](https://github.com/linq-team/linq-go/compare/v0.5.1...v0.6.0)

### Features

* make `from` optional on GET /v3/chats and add `to` filter ([8a556e9](https://github.com/linq-team/linq-go/commit/8a556e9bfe2be480ea168f6fe9525bae56d222ea))
* PDEV(Synapse): support markdown and text effects ([928c0b5](https://github.com/linq-team/linq-go/commit/928c0b50d0b9e6a79c9317430f9414ab33f6d3c3))
* Plt 397 patch update contact card endpoint rename my cards endpoints ([175e2ac](https://github.com/linq-team/linq-go/commit/175e2acfe96e1f562a0579f38bf4cc5dcbbad569))


### Chores

* **internal:** minor cleanup ([5cc3cae](https://github.com/linq-team/linq-go/commit/5cc3cae1c002314ac7024d40ece45228a64f7a23))
* **internal:** use explicit returns ([88e43a9](https://github.com/linq-team/linq-go/commit/88e43a9416323fa258978fc181f27fb0a2d21f55))
* **internal:** use explicit returns in more places ([9106e56](https://github.com/linq-team/linq-go/commit/9106e5607168f5804e9047757d01291b273e001a))

## 0.5.1 (2026-03-10)

Full Changelog: [v0.5.0...v0.5.1](https://github.com/linq-team/linq-go/compare/v0.5.0...v0.5.1)

### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([b7b858f](https://github.com/linq-team/linq-go/commit/b7b858fab1383b330c87e412f43f219b33ba9f78))

## 0.5.0 (2026-03-07)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/linq-team/linq-go/compare/v0.4.0...v0.5.0)

### Features

* Programmatically update contact card ([f3d2e8b](https://github.com/linq-team/linq-go/commit/f3d2e8b13e9adc583723eb423de0efebde8355e2))


### Chores

* **internal:** codegen related update ([c6c0327](https://github.com/linq-team/linq-go/commit/c6c0327e9fce2c1046658038edeb29214729ab1b))

## 0.4.0 (2026-03-05)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/linq-team/linq-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** fix shared ([eec3ea8](https://github.com/linq-team/linq-go/commit/eec3ea8555cc2726c3e5d7273653e7a45ab11fc7))
* **api:** update shared types ([b607b9a](https://github.com/linq-team/linq-go/commit/b607b9aaa3b40198a64391e3740a1331d52955c5))

## 0.3.0 (2026-03-05)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/linq-team/linq-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** add new endpoint ([5b3366a](https://github.com/linq-team/linq-go/commit/5b3366aebde57537fd323db056e02560a43b5224))

## 0.2.0 (2026-03-05)

Full Changelog: [v0.1.2...v0.2.0](https://github.com/linq-team/linq-go/compare/v0.1.2...v0.2.0)

### Features

* Allow 100 presigned URL or uploaded attachments (URL + ID) in a message ([2b32bbd](https://github.com/linq-team/linq-go/commit/2b32bbdf33ee1928793d062bde5bbed222c6ae8c))
* **api:** manual updates ([9d205d3](https://github.com/linq-team/linq-go/commit/9d205d312eb6a7e1ad9f6d4ead1a200a828e1ee0))
* Plt 361 synapse support editing messages in v3 ([2b7c25b](https://github.com/linq-team/linq-go/commit/2b7c25bff5a81de8a14619afcbb56a37bc3930ff))


### Bug Fixes

* remove unused part-level idempotency_key from OpenAPI spec ([19d060d](https://github.com/linq-team/linq-go/commit/19d060d98438be63c51097dbd9041fd7433f6b81))


### Chores

* **internal:** codegen related update ([b3938fe](https://github.com/linq-team/linq-go/commit/b3938fe08bb14412c79bde8d97f03379a5c044e4))

## 0.1.2 (2026-02-25)

Full Changelog: [v0.1.1...v0.1.2](https://github.com/linq-team/linq-go/compare/v0.1.1...v0.1.2)

### Bug Fixes

* sendReaction OpenAPI spec returns 202 not 200 ([4cde326](https://github.com/linq-team/linq-go/commit/4cde326c08dcb72282819888d5a2cecf0c5b8f20))


### Chores

* **internal:** move custom custom `json` tags to `api` ([a55460c](https://github.com/linq-team/linq-go/commit/a55460c1548a84dad5dc977cb5ef5b50437f1aeb))

## 0.1.1 (2026-02-24)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/linq-team/linq-go/compare/v0.1.0...v0.1.1)

### Features

* **api:** add shared resource ([54b07b5](https://github.com/linq-team/linq-go/commit/54b07b55f1b51ba33bf77dc903bc6c3f5acf4b34))
* **api:** wtf ([321a98e](https://github.com/linq-team/linq-go/commit/321a98e90e8ab8e1c978868afcb5c9181a44171b))
* PLT: Include sticker details for iMessage tapback webhooks ([204ec0b](https://github.com/linq-team/linq-go/commit/204ec0b243870bf968065a2d9912a2fa8879331c))

## 0.1.0 (2026-02-24)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/linq-team/linq-go/compare/v0.0.1...v0.1.0)

### Features

* Add Stainless SDK build to OpenAPI workflow ([1b4e2ee](https://github.com/linq-team/linq-go/commit/1b4e2eeb2062cdb6a57fd23aa0a088b81c184257))
* **api:** fix service type ([c390ca3](https://github.com/linq-team/linq-go/commit/c390ca3ac7093230ba97b390da045217fdeb8a55))
* **api:** manual updates ([9d015de](https://github.com/linq-team/linq-go/commit/9d015de2039f4c6c4b2e363d7e396eeb776977f2))
* PDEV(Synapse): Synapse: Support iMessage and RCS capability checker endpoints ([c561c13](https://github.com/linq-team/linq-go/commit/c561c13f057dcf95a444d5038534aff0b04bc010))


### Chores

* configure new SDK language ([7cc3634](https://github.com/linq-team/linq-go/commit/7cc36341f62db670db32ef3764cf6bf375952933))
* update SDK settings ([274599f](https://github.com/linq-team/linq-go/commit/274599fdef785f0155d1a7a576eea58c5de26683))
