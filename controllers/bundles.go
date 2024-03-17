package controllers

import (
	"mangathorg/internal/middlewares"
)

var IndexHandlerGetBundle = middlewares.Join(indexHandlerGet, middlewares.Log, middlewares.UserCheck)
var IndexHandlerPutBundle = middlewares.Join(indexHandlerPut, middlewares.Log, middlewares.Guard)
var IndexHandlerDeleteBundle = middlewares.Join(indexHandlerDelete, middlewares.Log, middlewares.Guard)
var IndexHandlerNoMethBundle = middlewares.Join(indexHandlerNoMeth, middlewares.Log, middlewares.UserCheck)
var IndexHandlerOtherBundle = middlewares.Join(indexHandlerOther, middlewares.Log, middlewares.UserCheck)
var LoginHandlerGetBundle = middlewares.Join(loginHandlerGet, middlewares.Log, middlewares.OnlyVisitors)
var LoginHandlerPostBundle = middlewares.Join(loginHandlerPost, middlewares.Log, middlewares.OnlyVisitors)
var RegisterHandlerGetBundle = middlewares.Join(registerHandlerGet, middlewares.Log, middlewares.OnlyVisitors)
var RegisterHandlerPostBundle = middlewares.Join(registerHandlerPost, middlewares.Log, middlewares.OnlyVisitors)
var ForgotPasswordHandlerGetBundle = middlewares.Join(forgotPasswordHandlerGet, middlewares.Log, middlewares.OnlyVisitors)
var HomeHandlerGetBundle = middlewares.Join(homeHandlerGet, middlewares.Log, middlewares.Guard)
var LogHandlerGetBundle = middlewares.Join(logHandlerGet, middlewares.Log, middlewares.UserCheck)
var ConfirmHandlerGetBundle = middlewares.Join(confirmHandlerGet, middlewares.Log, middlewares.OnlyVisitors)
var LogoutHandlerGetBundle = middlewares.Join(logoutHandlerGet, middlewares.Log, middlewares.Guard)
var PrincipalHandlerGetBundle = middlewares.Join(principalHandlerGet, middlewares.Log, middlewares.UserCheck)
var MangaRequestHandlerGet = middlewares.Join(mangaHandlerGet, middlewares.Log, middlewares.UserCheck)
var TagsHandlerGetBundle = middlewares.Join(tagsHandlerGet, middlewares.Log, middlewares.UserCheck)
var CategoryHandlerGetBundle = middlewares.Join(categoryHandlerGet, middlewares.Log, middlewares.UserCheck)
var FeedRequestHandlerGetBundle = middlewares.Join(feedRequestHandlerGet, middlewares.Log, middlewares.UserCheck)
var ChapterHandlerGetBundle = middlewares.Join(chapterHandlerGet, middlewares.Log, middlewares.UserCheck)
var MangaWholeRequestHandlerGetBundle = middlewares.Join(mangaWholeRequestHandlerGet, middlewares.Log, middlewares.UserCheck)
var CoversHandlerGetBundle = middlewares.Join(coverHandlerGet, middlewares.UserCheck)
var ScanHandlerGetBundle = middlewares.Join(scanHandlerGet, middlewares.UserCheck)
