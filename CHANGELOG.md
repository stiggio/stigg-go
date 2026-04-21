# Changelog

## 0.1.0-alpha.29 (2026-04-21)

Full Changelog: [v0.1.0-alpha.28...v0.1.0-alpha.29](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.28...v0.1.0-alpha.29)

### Features

* **api:** add endDate/startDate parameters to event credits get usage ([30e7c7c](https://github.com/stiggio/stigg-go/commit/30e7c7c15452bfb64810b1495477420ce48dc546))

## 0.1.0-alpha.28 (2026-04-16)

Full Changelog: [v0.1.0-alpha.27...v0.1.0-alpha.28](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.27...v0.1.0-alpha.28)

### Bug Fixes

* **api:** remove includeInactiveSubscriptions from usage.history, add event queue enum values ([8035a8e](https://github.com/stiggio/stigg-go/commit/8035a8ee3f7a32fcfb3a7cd14ec0c1cffee5a6a5))


### Documentation

* add custom currency event types to event queue parameters ([1823a17](https://github.com/stiggio/stigg-go/commit/1823a178539dd9a57036b28449ec34a20c30f315))

## 0.1.0-alpha.27 (2026-04-13)

Full Changelog: [v0.1.0-alpha.26...v0.1.0-alpha.27](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.26...v0.1.0-alpha.27)

### Bug Fixes

* **api:** remove includeInactiveSubscriptions from usage history, add event queue enum ([d807b0a](https://github.com/stiggio/stigg-go/commit/d807b0a46a3cabfaba2442cafc2251056a16a3b9))

## 0.1.0-alpha.26 (2026-04-08)

Full Changelog: [v0.1.0-alpha.25...v0.1.0-alpha.26](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.25...v0.1.0-alpha.26)

### Features

* **api:** add event type to event queue, remove param from usage history ([4263500](https://github.com/stiggio/stigg-go/commit/4263500bf2b186bfe1e51f517fc501d09749f8d0))
* **api:** add include_inactive_subscriptions parameter to usage history method ([e754fb2](https://github.com/stiggio/stigg-go/commit/e754fb216422d52e24358e23bba271763db7dc60))

## 0.1.0-alpha.25 (2026-04-07)

Full Changelog: [v0.1.0-alpha.24...v0.1.0-alpha.25](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.24...v0.1.0-alpha.25)

### Features

* **api:** stainless.yml ([14f7a2c](https://github.com/stiggio/stigg-go/commit/14f7a2c132ea60e6f58b4e615c10e5b531a3fea0))

## 0.1.0-alpha.24 (2026-04-06)

Full Changelog: [v0.1.0-alpha.23...v0.1.0-alpha.24](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.23...v0.1.0-alpha.24)

### Features

* **api:** add customer integrations and event queues endpoints ([7e2261b](https://github.com/stiggio/stigg-go/commit/7e2261b5dd8ba200a2ba25cffc3fd907fd8b3120))

## 0.1.0-alpha.23 (2026-03-28)

Full Changelog: [v0.1.0-alpha.22...v0.1.0-alpha.23](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.22...v0.1.0-alpha.23)

### Features

* **internal:** support comma format in multipart form encoding ([a3da9bb](https://github.com/stiggio/stigg-go/commit/a3da9bb62c698829e1c72abe15079092925f6b2a))


### Bug Fixes

* prevent duplicate ? in query params ([1c830f2](https://github.com/stiggio/stigg-go/commit/1c830f20ddef40e71f377dfb45a92ac808f7f1c1))


### Chores

* **ci:** skip lint on metadata-only changes ([f91b883](https://github.com/stiggio/stigg-go/commit/f91b8830f6c3d8f90352cabf292e4c49b7abf79f))
* **ci:** support opting out of skipping builds on metadata-only commits ([b5984a5](https://github.com/stiggio/stigg-go/commit/b5984a599b0d30616745fdfe029bbde29ebc7300))
* **client:** fix multipart serialisation of Default() fields ([cf4447b](https://github.com/stiggio/stigg-go/commit/cf4447b7b6095cdc47a465ffc66aca4f35e61753))
* **internal:** support default value struct tag ([af11d33](https://github.com/stiggio/stigg-go/commit/af11d33a9909509941d9c3c50f5b96e259dd446a))
* **internal:** update gitignore ([9e866ff](https://github.com/stiggio/stigg-go/commit/9e866ff5fa86ae8e69bd6d3734af5aba965e775d))
* remove unnecessary error check for url parsing ([0dfe9b1](https://github.com/stiggio/stigg-go/commit/0dfe9b111403b683ddf796627c5a42ee468deb8d))
* update docs for api:"required" ([84c9647](https://github.com/stiggio/stigg-go/commit/84c9647bed8275a0e1245fbbc8dc53b366f116c7))

## 0.1.0-alpha.22 (2026-03-22)

Full Changelog: [v0.1.0-alpha.21...v0.1.0-alpha.22](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.21...v0.1.0-alpha.22)

### Features

* **api:** api update ([455f667](https://github.com/stiggio/stigg-go/commit/455f667bc4ad37af6ba0a8ab48fea31f35ef5393))
* **api:** api update ([0cb10fb](https://github.com/stiggio/stigg-go/commit/0cb10fbe13eee79e0d13dc062cb5256e2e546e69))
* **api:** api update ([01a32c9](https://github.com/stiggio/stigg-go/commit/01a32c966f1cba885060dbbdfdb0a1efbd9935e4))
* **api:** api update ([fc7ff54](https://github.com/stiggio/stigg-go/commit/fc7ff54f8c49606fdb2c20e12e5652799483df9a))
* **api:** api update ([f27c558](https://github.com/stiggio/stigg-go/commit/f27c558819b2676c1334f6aa7d5c4467f2316553))

## 0.1.0-alpha.21 (2026-03-17)

Full Changelog: [v0.1.0-alpha.20...v0.1.0-alpha.21](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.20...v0.1.0-alpha.21)

### Features

* **api:** api update ([c11adee](https://github.com/stiggio/stigg-go/commit/c11adee2d440aca786ffdf72660dee054f72bdb9))
* **api:** api update ([a127b7c](https://github.com/stiggio/stigg-go/commit/a127b7c972d9ca4c6c8257e84594373e3fe7bf78))
* **api:** updated stainless config with new endpoint ([a95bf2d](https://github.com/stiggio/stigg-go/commit/a95bf2d4c359832cb2cbfac0ca1455826ea99dfc))

## 0.1.0-alpha.20 (2026-03-17)

Full Changelog: [v0.1.0-alpha.19...v0.1.0-alpha.20](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.19...v0.1.0-alpha.20)

### Features

* **api:** api update ([9d0d3a0](https://github.com/stiggio/stigg-go/commit/9d0d3a09b1f91f37924c6ccfde2e8fd9a80c5b24))
* **api:** api update ([28ed5f2](https://github.com/stiggio/stigg-go/commit/28ed5f20fbce2109b18751cc7ddaea9273997644))


### Chores

* **internal:** tweak CI branches ([0d2ba59](https://github.com/stiggio/stigg-go/commit/0d2ba598eb2a914a94719011a97cfbe65b324c9d))

## 0.1.0-alpha.19 (2026-03-16)

Full Changelog: [v0.1.0-alpha.18...v0.1.0-alpha.19](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.18...v0.1.0-alpha.19)

### Features

* **api:** api update ([c866c60](https://github.com/stiggio/stigg-go/commit/c866c60890807e0e170fbb4f452eff6a5a12e39c))
* **api:** api update ([053445b](https://github.com/stiggio/stigg-go/commit/053445bdf80a76087eb765e2376966637a217516))
* **api:** api update ([468091c](https://github.com/stiggio/stigg-go/commit/468091c3b7e0c27fc6d63710693fc0d701f482bf))

## 0.1.0-alpha.18 (2026-03-12)

Full Changelog: [v0.1.0-alpha.17...v0.1.0-alpha.18](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.17...v0.1.0-alpha.18)

### Features

* **api:** api update ([b028427](https://github.com/stiggio/stigg-go/commit/b028427e5f59035c5686a0a12172018ea3f3c798))


### Chores

* **internal:** minor cleanup ([9dbe222](https://github.com/stiggio/stigg-go/commit/9dbe2224a424aaafbbe566c30c17ca99c7c7ceb1))
* **internal:** use explicit returns ([cdb4afb](https://github.com/stiggio/stigg-go/commit/cdb4afbccf16c1c049eab41ea69eeb3f64c24283))
* **internal:** use explicit returns in more places ([a7ab8a6](https://github.com/stiggio/stigg-go/commit/a7ab8a608f213b9ffa520e1f02233d9a26a068c0))

## 0.1.0-alpha.17 (2026-03-10)

Full Changelog: [v0.1.0-alpha.16...v0.1.0-alpha.17](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.16...v0.1.0-alpha.17)

### Features

* **api:** added credits resources ([0f3ca7e](https://github.com/stiggio/stigg-go/commit/0f3ca7e46e5c044cc6ad59dc5719182942264d9a))
* **api:** api update ([7ccb79a](https://github.com/stiggio/stigg-go/commit/7ccb79a52938f9e918a097ba345a926b73041f90))
* **api:** api update ([3b8cbcc](https://github.com/stiggio/stigg-go/commit/3b8cbcc7ae5e726647fcf9c4ded02873c8613fd5))
* **api:** api update ([213a89f](https://github.com/stiggio/stigg-go/commit/213a89ff9ca61a6c40951d6c3a1a65cce35389f1))
* **api:** api update ([245dccb](https://github.com/stiggio/stigg-go/commit/245dccbfd8293bcfecf7534dda4037ea8776baed))

## 0.1.0-alpha.16 (2026-03-08)

Full Changelog: [v0.1.0-alpha.15...v0.1.0-alpha.16](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.15...v0.1.0-alpha.16)

### Features

* **api:** api update ([e4318e0](https://github.com/stiggio/stigg-go/commit/e4318e08ad54df9334a07c005daf93bec43c18ab))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([89c72f6](https://github.com/stiggio/stigg-go/commit/89c72f69b0a64cab7d976b104eaa8f2e6c7b99f9))
* **internal:** codegen related update ([5a443f6](https://github.com/stiggio/stigg-go/commit/5a443f67ca4a53a1fec8fd61f2652fd470a7b48f))

## 0.1.0-alpha.15 (2026-03-05)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Features

* **api:** api update ([4dbb21b](https://github.com/stiggio/stigg-go/commit/4dbb21ba45a90c969101603479de2641b90ba8f6))

## 0.1.0-alpha.14 (2026-03-05)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Features

* **api:** api update ([7d0ad48](https://github.com/stiggio/stigg-go/commit/7d0ad48e77ca14ac9d66e7c7df259beab4bda5a8))

## 0.1.0-alpha.13 (2026-03-04)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Features

* **api:** api update ([85be18f](https://github.com/stiggio/stigg-go/commit/85be18f610d862e3416443064ccbdd98272fb8d6))


### Chores

* **internal:** codegen related update ([0646161](https://github.com/stiggio/stigg-go/commit/06461616ca2be20f115a5ba6bc7d755d8a328343))

## 0.1.0-alpha.12 (2026-03-02)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Features

* **api:** api update ([743d67a](https://github.com/stiggio/stigg-go/commit/743d67a80e9f0cf75f977be1d2c5f39b7e901a78))

## 0.1.0-alpha.11 (2026-03-02)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Features

* **api:** api update ([666668c](https://github.com/stiggio/stigg-go/commit/666668c38806537b067927265c481fa5cae3949e))

## 0.1.0-alpha.10 (2026-02-27)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Features

* **api:** api update ([4c17d93](https://github.com/stiggio/stigg-go/commit/4c17d93cb8e7f3c957a62d636fe8b7ce8f6b965a))
* **api:** stainless config updates ([c68b5fd](https://github.com/stiggio/stigg-go/commit/c68b5fd085f034c598d724860f9b36a96511c1bf))
* **api:** update endpoints and models ([876942a](https://github.com/stiggio/stigg-go/commit/876942a17990cc2b24b5ed1a706d736698bcc440))


### Chores

* **internal:** move custom custom `json` tags to `api` ([1598cbd](https://github.com/stiggio/stigg-go/commit/1598cbd538c4e638358b85f2aa5f4d4c00c4a514))

## 0.1.0-alpha.9 (2026-02-22)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* **api:** api update ([b9f5caa](https://github.com/stiggio/stigg-go/commit/b9f5caaecdf0712b3d428d310a2e2769e6158738))

## 0.1.0-alpha.8 (2026-02-22)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Bug Fixes

* allow canceling a request while it is waiting to retry ([0fe754d](https://github.com/stiggio/stigg-go/commit/0fe754d84b188c647fdbd51412bc524a084e5311))
* **internal:** skip tests that depend on mock server ([ca0e8c3](https://github.com/stiggio/stigg-go/commit/ca0e8c39ab3c4dad50603b29d019bb778c7abf9e))


### Chores

* **internal:** remove mock server code ([5cdeae5](https://github.com/stiggio/stigg-go/commit/5cdeae5cfe1ee7807cae9d947bae2c7bf9e2c614))
* update mock server docs ([87e1809](https://github.com/stiggio/stigg-go/commit/87e18090b0217a9c94801608d9be1dc2bf3ef48e))

## 0.1.0-alpha.7 (2026-02-18)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** add additional endpoints ([c52ee55](https://github.com/stiggio/stigg-go/commit/c52ee5563ad37eab8d0b9b01ffb0e822d76fb8e3))
* **api:** add features/addons/plans resources, subscription usage/invoice, products/coupons methods ([18b6a61](https://github.com/stiggio/stigg-go/commit/18b6a6116882a7327f7888372769b653032ef19e))
* **api:** Add missing endpoints ([353159c](https://github.com/stiggio/stigg-go/commit/353159c31e95542b913d16e5670a7d9f1adbee39))
* **api:** api update ([db8b99a](https://github.com/stiggio/stigg-go/commit/db8b99ab7036ea8ba07ac0ba11bbfd548606ce4a))
* **api:** api update ([5957eb2](https://github.com/stiggio/stigg-go/commit/5957eb2201978aa3aec72190e54000ef048367e9))
* **api:** api update ([b203a1c](https://github.com/stiggio/stigg-go/commit/b203a1c5be3a8f35dee993f4c586e695191ee9b2))
* **api:** api update ([21a0d84](https://github.com/stiggio/stigg-go/commit/21a0d84e48f99072260b9b6cdf4615a187b22692))
* **api:** api update ([9cd9c4d](https://github.com/stiggio/stigg-go/commit/9cd9c4d31ecd38b8c7983ec61a31ddb85a239380))
* **api:** api update ([0a69fa7](https://github.com/stiggio/stigg-go/commit/0a69fa7517866edcd1a3202ee3d3223193b11743))
* **api:** api update ([848f506](https://github.com/stiggio/stigg-go/commit/848f5069110890d3832cb5f65993c6faf20b67fa))
* **api:** api update ([7ca40ec](https://github.com/stiggio/stigg-go/commit/7ca40ecdbb3af15330b42c4d59e539ef615a405b))
* **api:** api update ([d3b9008](https://github.com/stiggio/stigg-go/commit/d3b9008e12c45a4fec58e06525e19bb4cfb567fe))
* **api:** api update ([6177c47](https://github.com/stiggio/stigg-go/commit/6177c478a4f88a26f28008a46bf02e715fdf8125))
* **api:** manual updates ([a13d3d0](https://github.com/stiggio/stigg-go/commit/a13d3d0fc07c1a2a2488b69c9946d1ec661f5a38))
* **api:** manual updates ([a13d3d0](https://github.com/stiggio/stigg-go/commit/a13d3d0fc07c1a2a2488b69c9946d1ec661f5a38))
* **api:** manual updates ([b4fdfc4](https://github.com/stiggio/stigg-go/commit/b4fdfc47422a1ce6f6d14fef525c095df4d48248))
* **api:** manual updates ([0808829](https://github.com/stiggio/stigg-go/commit/08088297d51ec287be96bb252244aa8720f66c09))
* **api:** manual updates ([a44cc9c](https://github.com/stiggio/stigg-go/commit/a44cc9c9fa5b99f21ed87965c476fe9a0d9e1903))
* **api:** trigger release ([cfa32cc](https://github.com/stiggio/stigg-go/commit/cfa32cc63a71f284ae9064dbf6a6fdef69099d21))
* **api:** trigger release ([cfa32cc](https://github.com/stiggio/stigg-go/commit/cfa32cc63a71f284ae9064dbf6a6fdef69099d21))
* **api:** updated the production environment ([06e6e51](https://github.com/stiggio/stigg-go/commit/06e6e51ee7546d2a998f0df5707a120ed0773010))


### Bug Fixes

* **encoder:** correctly serialize NullStruct ([815ebf3](https://github.com/stiggio/stigg-go/commit/815ebf31b26fa904cf3071d3358890a9b9bfa2fd))


### Chores

* configure new SDK language ([91f2a31](https://github.com/stiggio/stigg-go/commit/91f2a310c07bd02c05c30783f3b6ea7c2b903dc6))

## 0.1.0-alpha.6 (2026-02-08)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** api update ([1c823bd](https://github.com/stiggio/stigg-go/commit/1c823bd05288d41495fe2295b69c165d56719551))
* **api:** api update ([be245bb](https://github.com/stiggio/stigg-go/commit/be245bb99ead798783076a5019c5d4c1a72145b9))
* **api:** api update ([c994cc8](https://github.com/stiggio/stigg-go/commit/c994cc861dc31bb35fe0f168c9dc3131deb2a14e))
* **api:** manual updates ([66789dc](https://github.com/stiggio/stigg-go/commit/66789dc58bc0600b7c07bcf8df96ed5871349f1c))
* **api:** manual updates ([8027a42](https://github.com/stiggio/stigg-go/commit/8027a42946d045818d694bfb306568a79339fe47))

## 0.1.0-alpha.5 (2026-01-29)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** update stainless config ([143ef26](https://github.com/stiggio/stigg-go/commit/143ef2615a0112ff8471c5bc5db3a75b95138864))

## 0.1.0-alpha.4 (2026-01-28)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** api update ([d26cad9](https://github.com/stiggio/stigg-go/commit/d26cad994934ee14173e2579a321e4b033903ba9))
* **api:** api update ([2e0a3a1](https://github.com/stiggio/stigg-go/commit/2e0a3a1ddefb6122b07aec7f0db719288e92d08e))

## 0.1.0-alpha.3 (2026-01-28)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** api update ([9c94833](https://github.com/stiggio/stigg-go/commit/9c94833e2e97bbab05a3940f6a8cc382c721a9ee))
* **api:** api update ([4326c32](https://github.com/stiggio/stigg-go/commit/4326c32389a45b94eebcaad7daf1d73149fd81fe))
* **api:** improved cursor pagination ([116b768](https://github.com/stiggio/stigg-go/commit/116b768e6c87ad5d64892b0545ead907f38164e8))

## 0.1.0-alpha.2 (2026-01-27)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/stiggio/stigg-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** api update ([4422d29](https://github.com/stiggio/stigg-go/commit/4422d2979239c8905d59cdd57221c64f229edf39))
* **api:** comment out promotional endpoints ([c1b5d46](https://github.com/stiggio/stigg-go/commit/c1b5d462419b850012f4488fa104aa9416cfeaf2))

## 0.1.0-alpha.1 (2026-01-26)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/stiggio/stigg-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** added customer endpoints ([a726f1f](https://github.com/stiggio/stigg-go/commit/a726f1feb56902ed664f4766acd3995caded809f))
* **api:** api update ([976e691](https://github.com/stiggio/stigg-go/commit/976e69110c9596d37c19d397497f8b706c4f801c))
* **api:** api update ([14e6ec8](https://github.com/stiggio/stigg-go/commit/14e6ec87475aecc528481cf9cf8f2bb28678653f))
* **api:** api update ([b7d6395](https://github.com/stiggio/stigg-go/commit/b7d6395fe58082f207da27ac5ac38e1288aa0a6a))
* **api:** manual updates ([d84ff98](https://github.com/stiggio/stigg-go/commit/d84ff982ee4888e4650e82910dfafa5e22428938))
* **api:** manual updates ([f86edb9](https://github.com/stiggio/stigg-go/commit/f86edb9ac097f01397050235a689dd38e307d3f4))
* **api:** manual updates ([126bfbc](https://github.com/stiggio/stigg-go/commit/126bfbc0e5587a9895a5804de82fb02addac00a2))
* **api:** manual updates ([5c7c363](https://github.com/stiggio/stigg-go/commit/5c7c363218b2c18673aa14b790f6d5696c1845c8))
* **api:** update via SDK Studio ([fd0e5b9](https://github.com/stiggio/stigg-go/commit/fd0e5b9ca7fb1f7a7d4a0373def71547bb6b69b5))
* **api:** update via SDK Studio ([db6f3f5](https://github.com/stiggio/stigg-go/commit/db6f3f5b4c3d4a6d112f57b3a25610f5e222ead9))
* **api:** update via SDK Studio ([8dddf21](https://github.com/stiggio/stigg-go/commit/8dddf21af09f21083f0db49e86cec7b4a900de94))
* **client:** add a convenient param.SetJSON helper ([6792b52](https://github.com/stiggio/stigg-go/commit/6792b52e9a0b98180239a5e7632321c3baa661c1))
* **client:** support optional json html escaping ([e0a9b89](https://github.com/stiggio/stigg-go/commit/e0a9b8900d66b71682171c0b95f4b29f81ee0b75))


### Bug Fixes

* **client:** process custom base url ahead of time ([9de2d19](https://github.com/stiggio/stigg-go/commit/9de2d190611141438439869d21fad9c26a2e4517))
* **docs:** add missing pointer prefix to api.md return types ([233fbc3](https://github.com/stiggio/stigg-go/commit/233fbc35b23b5266e96c3eeed262bc77df5fe67b))
* skip usage tests that don't work with Prism ([71d3349](https://github.com/stiggio/stigg-go/commit/71d3349d6ac9d73ca03358d183799de7e8399d83))


### Chores

* add float64 to valid types for RegisterFieldValidator ([fe3f124](https://github.com/stiggio/stigg-go/commit/fe3f12446dc61cf2ce6327fc27b202a58ac9d7e7))
* bump gjson version ([a771236](https://github.com/stiggio/stigg-go/commit/a771236ce8d695732b72602cfe670642aca61aff))
* configure new SDK language ([8dea2f4](https://github.com/stiggio/stigg-go/commit/8dea2f48c4860af7e122982cc041336b66d24318))
* **internal:** codegen related update ([ba6727c](https://github.com/stiggio/stigg-go/commit/ba6727c16b0a05713fe8c77823072ef5f0f8cb28))
* **internal:** codegen related update ([b3fccd2](https://github.com/stiggio/stigg-go/commit/b3fccd277042577929dcfd655c8fe0856cab1ae3))
* **internal:** fix lint script for tests ([3c041fa](https://github.com/stiggio/stigg-go/commit/3c041fa9f5fc78fcbdb0b8e76c4f6d926a85b136))
* **internal:** grammar fix (it's -&gt; its) ([a96952f](https://github.com/stiggio/stigg-go/commit/a96952f3d9947a15cec880659b0b8f8093d02c9c))
* **internal:** update `actions/checkout` version ([bf8e893](https://github.com/stiggio/stigg-go/commit/bf8e8934f62d48c9131f698b027c8c69eeb18e19))
* **internal:** update comment in script ([d8da697](https://github.com/stiggio/stigg-go/commit/d8da69788580ac0fa8dbda8e14d14a7ba4638d8a))
* lint tests ([9ae23d3](https://github.com/stiggio/stigg-go/commit/9ae23d3bbd63362059b1cfa121f0fdf0762694ff))
* lint tests in subpackages ([15428d9](https://github.com/stiggio/stigg-go/commit/15428d92447da679b44c98a5509ed70d51fd94e5))
* update @stainless-api/prism-cli to v5.15.0 ([29db473](https://github.com/stiggio/stigg-go/commit/29db473fd503a4ffcb238f7324f019cb695dcce9))
* update SDK settings ([7e6de05](https://github.com/stiggio/stigg-go/commit/7e6de052461f936105a269e99c30c1b47a05f18d))
