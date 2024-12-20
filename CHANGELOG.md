# Changelog

## [1.9.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.8.2...v1.9.0) (2024-12-05)


### Features

* **contact:** add npsScore attribute ([4f8b610](https://www.github.com/Karnott/skalin-sdk/commit/4f8b610dbf0db2911e3e85ace11003671d433b0b))

### [1.8.2](https://www.github.com/Karnott/skalin-sdk/compare/v1.8.1...v1.8.2) (2024-03-27)


### Bug Fixes

* **getEntities:** use 'hasNextPage' attribute to get all entities ([d5ceca7](https://www.github.com/Karnott/skalin-sdk/commit/d5ceca7cffc227049e9f66961ceb57d0f8adf29e))

### [1.8.1](https://www.github.com/Karnott/skalin-sdk/compare/v1.8.0...v1.8.1) (2024-02-28)


### Bug Fixes

* **skalin:** add DeleteContact func to the Skalin interface definition ([1119d17](https://www.github.com/Karnott/skalin-sdk/commit/1119d172ab5df1b6c6ebf41b7c20ea9fbff54d35))

## [1.8.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.7.0...v1.8.0) (2024-02-28)


### Features

* **agreement:** delete agreement ([5b2eddc](https://www.github.com/Karnott/skalin-sdk/commit/5b2eddc0d6279c272efc47a5e9c3466e799a328f))
* **contact:** add DeleteContact func ([186bff7](https://www.github.com/Karnott/skalin-sdk/commit/186bff7bba284cf53286aee9c7861c51fc333301))

## [1.7.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.6.0...v1.7.0) (2023-10-18)


### Features

* **tags:** create GetTags and GetTagByID func ([8038241](https://www.github.com/Karnott/skalin-sdk/commit/80382411857d03933e748228932a73fc5b178163))

## [1.6.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.5.0...v1.6.0) (2023-07-19)


### Features

* **tracking:** add CIP in HitTrack struct ([b0ba4eb](https://www.github.com/Karnott/skalin-sdk/commit/b0ba4ebfcd7ad172d2aca320c4be809259660d5b))

## [1.5.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.4.0...v1.5.0) (2023-07-17)


### Features

* add custom headers for hit endpoint ([01f131e](https://www.github.com/Karnott/skalin-sdk/commit/01f131e73bbb501ccc24915c08d366485cdcefcc))

## [1.4.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.3.0...v1.4.0) (2023-07-12)


### Features

* add tracking endpoint ([e90ea45](https://www.github.com/Karnott/skalin-sdk/commit/e90ea45d083c6a8eb1e419144bc5b5cc2fb708d6))

## [1.3.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.2.0...v1.3.0) (2023-06-21)


### Features

* **getData:** update generic getEntities func to get all data using metadata returning by skalin API ([dfbf5c4](https://www.github.com/Karnott/skalin-sdk/commit/dfbf5c4a220a02f3f2c7df206a57d33d73fcfbde))

## [1.2.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.1.2...v1.2.0) (2023-06-19)


### Features

* **agreement:** create sdk func for agreement ([3f61bf2](https://www.github.com/Karnott/skalin-sdk/commit/3f61bf2fc9d824e5186af10abff1eac6f487a7c5))

### [1.1.2](https://www.github.com/Karnott/skalin-sdk/compare/v1.1.1...v1.1.2) (2023-06-15)


### Bug Fixes

* **api:** log body response error fix + do not log body if it contains access_token ([394467d](https://www.github.com/Karnott/skalin-sdk/commit/394467d2221c229780b92c14e6ce049104959f9a))

### [1.1.1](https://www.github.com/Karnott/skalin-sdk/compare/v1.1.0...v1.1.1) (2023-06-08)


### Bug Fixes

* **logger:** define default logger in new func ([a2c6a5b](https://www.github.com/Karnott/skalin-sdk/commit/a2c6a5bab962c0103f5e8f38cec5e39bccfb19f0))

## [1.1.0](https://www.github.com/Karnott/skalin-sdk/compare/v1.0.0...v1.1.0) (2023-06-07)


### Features

* **contact:** consider params to filter contacts list ([64aa4bf](https://www.github.com/Karnott/skalin-sdk/commit/64aa4bf1203d3804f8ac7feb019515260dbf5c4d))
* **contact:** create contact with customerId url + consider customAttributes ([73fe116](https://www.github.com/Karnott/skalin-sdk/commit/73fe116eb48641e87b04452e16492fdf2b2b6d55))
* **contact:** create update func ([c3ffa11](https://www.github.com/Karnott/skalin-sdk/commit/c3ffa1115f973d6b6177b4ce9d667a16c9092aba))
* **customer:** create getCustomers func ([b1456e7](https://www.github.com/Karnott/skalin-sdk/commit/b1456e720ce6413c9b3ff75a56918ea89dc1b54d))
* **skalin:** add setter for logger ([a5ff6bf](https://www.github.com/Karnott/skalin-sdk/commit/a5ff6bf4c1f2b86fd870fbd3508c95274c1a7841))


### Bug Fixes

* **updateContact:** remove unmarshal response from update API request because skalin API only return 'status: success' response ([1c3e032](https://www.github.com/Karnott/skalin-sdk/commit/1c3e03246bba682127782c84bb01f5c80098bec6))

## 1.0.0 (2023-06-02)


### Features

* create sdk function to create contact and customer + list contact ([0d21014](https://www.github.com/Karnott/skalin-sdk/commit/0d21014fae63718ca8acf153c2e8e6a0b84055b6))
