/*-----------------------------------------------------------------------
 *                     错误代码表
 *
 *-----------------------------------------------------------------------
 */

#ifndef	__ERRORMSG_H__
#define __ERRORMSG_H__
#include "StdAfx.h"

#define	ERR_EXCEPTION	-9999			//截获异常

#define ERR_OK			0				//交易成功

#define ERR_VER			-1				//版本不符
#define ERR_RETCODE		-2				//返回码不对
#define ERR_LENGTH		-3				//数据长度不对
#define ERR_FILENAME	-4				//文件名非法
#define ERR_FILESTAT	-5				//文件访问状态非法
#define ERR_FAIL		-6				//操作失败

/*sios error : ERR_SIOS_ */
#define	ERR_SIOS_NOREC		-100		//指定的记录不存在
#define ERR_SIOS_DOWNLOAD	-101		//下载文件失败


/*net error : ERR_NET_ */
#define ERR_NET_CONNECT		-200		//网络连接不通
#define ERR_NET_SEND		-201		//数据发送出错
#define ERR_NET_RECV		-202		//数据接收出错
#define ERR_NET_RECVFILE	-203		//接收文件出错
#define ERR_NET_SENDFILE	-204		//发送文件出错

/*trans errot : ERR_TRN_ */
#define ERR_TRN_SUBCODE		-300		//无效的子系统代码
#define ERR_TRN_STATION		-301		//无效的站点号
#define ERR_TRN_SIGNPWD		-302		//无效的签到密码


/*EN_CARD error : ERR_ENCARD_ */
#define ERR_ENCARD_RHEAD	-500		//读加密卡头错
#define ERR_ENCARD_CONFIG	-501		//读配置区出错
#define ERR_ENCARD_RKEY		-502		//读密钥错
#define ERR_ENCARD_OPEN		-503		//打开加密卡失败
#define ERR_PSAM_INIT		-504		//psam卡初始化失败
#define ERR_PSAM_FACTORY	-505		//读厂家信息失败
#define ERR_PSAM_LOCAL		-506		//读本地配置信息失败
#define ERR_PSAM_SERVER		-507		//读服务器配置信息失败
#define ERR_PSAM_KEY		-508		//读密钥配置信息失败
#define ERR_PSAM_AUTH		-509		//读厂家授权信息失败
#define ERR_PSAM_SAUTH		-510		//读子系统授权信息失败


/*DLL error : ERR_DLL_ from -1000 */
#define ERR_DLL_SIOS		-1001		//SIOS没有正常运行
#define ERR_DLL_DSQL		-1002		//DSQL操作错误
#define ERR_DLL_BUF_MIN		-1003		//分配的缓冲区太小，不能拷贝
#define ERR_DLL_UNPACK		-1004		//解包出错
#define ERR_DLL_REDO		-1005		//重做业务2003-09-05
#define ERR_DLL_NOPHOTO		-1006		//没有相片文件
#define ERR_DLL_NOFILE		-1007		//指定文件不存在

/*定义升级返回值 from 1100*/
#define	ERR_FILEEXIST		-1100		//文件已经存在
#define	ERR_REFUSE			-1101		//操作被拒绝
#define	ERR_NO_FILE			-1102		//没有文件
#define	ERR_DEL_FAIL		-1103		//删除文件失败
#define	ERR_COMM_FAIL		-1104		//通讯失败


/*第三方返回值定义 from 1200*/
#define ERR_TA_TRANAMT		-1200		//交易额错误
#define ERR_TA_NOT_INIT		-1201		//第三方API没有初始化
#define ERR_TA_CARDREADER	-1202		//读卡器错误
#define ERR_TA_READCARD		-1203		//读卡失败
#define ERR_TA_WRITECARD	-1204		//写卡失败
#define ERR_TA_LIMIT_FUNC	-1205		//函数调用功能限制
#define ERR_TA_CARDTYPE		-1206		//不是消费卡
#define ERR_TA_SNO			-1207		//非本院校卡
#define ERR_TA_EXPIRECARD	-1208		//过期卡
#define ERR_TA_FAIL_CHGUT	-1209		//修改用卡次数失败
#define ERR_TA_NOT_SAMECARD	-1210		//写卡时卡号不符
#define ERR_TA_WRONG_PWD	-1211		//卡消费时输入密码错误
#define ERR_TA_LOW_BALAN	-1212		//卡内余额不足
#define ERR_TA_EXCEED_QUOTA	-1213		//超过消费限额
#define ERR_TA_LOST_CARD	-1214		//挂失卡
#define ERR_TA_FREEZE_CARD	-1215		//冻结卡
#define ERR_TA_CARDNO		-1216		//卡号帐号不符
#define ERR_TA_ID_CLOSE		-1217		//身份关闭
#define ERR_TA_CR_DLL		-1218		//加载读卡器动态链接库失败
#define ERR_TA_CR_INIT		-1219		//读卡器初始化失败
#define ERR_TA_PARA			-1220		//参数错误
#define ERR_TA_NOREC		-1221		//没有这个帐户
#define ERR_TA_SUB_SUCC		-1222		//补助成功,也是正确的返回信息
#define ERR_TA_SUB_FAIL		-1223		//补助失败,也是正确的返回信息
#define ERR_TA_INITED		-1224		//读卡器已经初始化，请关闭
#define ERR_TA_UP_DLL		-1225		//加载UpdateAPI.dll动态库失败
#define ERR_TA_WRITECARD2	-1226		//回写卡余额失败
#define ERR_TA_QUERYACC		-1227		//查询帐户失败
#define ERR_TA_REREADCARD	-1228		//写卡时候重新读卡余额失败
#define ERR_TA_NO_FUN		-1229		//没有这个函数

/*定义设置同步服务器返回值 from -1300*/
#define ERR_SS_MAX_NG		-1300		//增加同步网关时超过最多数目
#define ERR_SS_NO_NG		-1301		//删除同步网关时没有找到网关
#define ERR_SS_DOWN_TSK		-1302		//增加同步网关时下载同步任务失败
#define ERR_SS_UP_TSK		-1303		//删除同步网关时上传同步任务失败

#endif