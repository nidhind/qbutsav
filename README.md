# QBUtsav
Server for QBUtsav Auction

### Environment variables to export

| Name | Description | Default Value |
| ------ | ------ | ------ |
|`GO_ENV`|Running environment for APP. Possible values are `development`, `production`|`development`|
|`PORT`|`port` on which the server will listen and serve|`8080`|
|`GIN_LOG_BASE_PATH`|Absolute path to the folder in which the `gin log` file must be created.WARNING: If not set then tmp folder will be used and the data will be lost when the instance restarts on normal systems|`/tmp`|
|`GIN_LOG_FIlE_NAME`|Filename to be used for the `gin log` file|`gin.log`|
