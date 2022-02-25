1. 判断逻辑，先进行本地的计算，cache的计算。
2. 然后再发http请求，http请求一般都比较慢。
3. 增加保护功能不应该报错，而是做降级处理， 不报错,但是要发警告。

``` js
    var response = SomeService.getsomestate(self, getsomestateParams, GLOBAL.timeout);
    if (!response || response.RetCode !== 0) {
        self.logger.info("get CheckWTFCreationAuthStatus failed:", response.RetCode);
        ret = Func.generateErrorCode("ACTION_ERROR", "SomeService.getsomestate failed");
        task.emit("check_university_auth_status_finish", ret);
        return;
    }
```