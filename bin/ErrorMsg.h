/*-----------------------------------------------------------------------
 *                     ��������
 *
 *-----------------------------------------------------------------------
 */

#ifndef	__ERRORMSG_H__
#define __ERRORMSG_H__
#include "StdAfx.h"

#define	ERR_EXCEPTION	-9999			//�ػ��쳣

#define ERR_OK			0				//���׳ɹ�

#define ERR_VER			-1				//�汾����
#define ERR_RETCODE		-2				//�����벻��
#define ERR_LENGTH		-3				//���ݳ��Ȳ���
#define ERR_FILENAME	-4				//�ļ����Ƿ�
#define ERR_FILESTAT	-5				//�ļ�����״̬�Ƿ�
#define ERR_FAIL		-6				//����ʧ��

/*sios error : ERR_SIOS_ */
#define	ERR_SIOS_NOREC		-100		//ָ���ļ�¼������
#define ERR_SIOS_DOWNLOAD	-101		//�����ļ�ʧ��


/*net error : ERR_NET_ */
#define ERR_NET_CONNECT		-200		//�������Ӳ�ͨ
#define ERR_NET_SEND		-201		//���ݷ��ͳ���
#define ERR_NET_RECV		-202		//���ݽ��ճ���
#define ERR_NET_RECVFILE	-203		//�����ļ�����
#define ERR_NET_SENDFILE	-204		//�����ļ�����

/*trans errot : ERR_TRN_ */
#define ERR_TRN_SUBCODE		-300		//��Ч����ϵͳ����
#define ERR_TRN_STATION		-301		//��Ч��վ���
#define ERR_TRN_SIGNPWD		-302		//��Ч��ǩ������


/*EN_CARD error : ERR_ENCARD_ */
#define ERR_ENCARD_RHEAD	-500		//�����ܿ�ͷ��
#define ERR_ENCARD_CONFIG	-501		//������������
#define ERR_ENCARD_RKEY		-502		//����Կ��
#define ERR_ENCARD_OPEN		-503		//�򿪼��ܿ�ʧ��
#define ERR_PSAM_INIT		-504		//psam����ʼ��ʧ��
#define ERR_PSAM_FACTORY	-505		//��������Ϣʧ��
#define ERR_PSAM_LOCAL		-506		//������������Ϣʧ��
#define ERR_PSAM_SERVER		-507		//��������������Ϣʧ��
#define ERR_PSAM_KEY		-508		//����Կ������Ϣʧ��
#define ERR_PSAM_AUTH		-509		//��������Ȩ��Ϣʧ��
#define ERR_PSAM_SAUTH		-510		//����ϵͳ��Ȩ��Ϣʧ��


/*DLL error : ERR_DLL_ from -1000 */
#define ERR_DLL_SIOS		-1001		//SIOSû����������
#define ERR_DLL_DSQL		-1002		//DSQL��������
#define ERR_DLL_BUF_MIN		-1003		//����Ļ�����̫С�����ܿ���
#define ERR_DLL_UNPACK		-1004		//�������
#define ERR_DLL_REDO		-1005		//����ҵ��2003-09-05
#define ERR_DLL_NOPHOTO		-1006		//û����Ƭ�ļ�
#define ERR_DLL_NOFILE		-1007		//ָ���ļ�������

/*������������ֵ from 1100*/
#define	ERR_FILEEXIST		-1100		//�ļ��Ѿ�����
#define	ERR_REFUSE			-1101		//�������ܾ�
#define	ERR_NO_FILE			-1102		//û���ļ�
#define	ERR_DEL_FAIL		-1103		//ɾ���ļ�ʧ��
#define	ERR_COMM_FAIL		-1104		//ͨѶʧ��


/*����������ֵ���� from 1200*/
#define ERR_TA_TRANAMT		-1200		//���׶����
#define ERR_TA_NOT_INIT		-1201		//������APIû�г�ʼ��
#define ERR_TA_CARDREADER	-1202		//����������
#define ERR_TA_READCARD		-1203		//����ʧ��
#define ERR_TA_WRITECARD	-1204		//д��ʧ��
#define ERR_TA_LIMIT_FUNC	-1205		//�������ù�������
#define ERR_TA_CARDTYPE		-1206		//�������ѿ�
#define ERR_TA_SNO			-1207		//�Ǳ�ԺУ��
#define ERR_TA_EXPIRECARD	-1208		//���ڿ�
#define ERR_TA_FAIL_CHGUT	-1209		//�޸��ÿ�����ʧ��
#define ERR_TA_NOT_SAMECARD	-1210		//д��ʱ���Ų���
#define ERR_TA_WRONG_PWD	-1211		//������ʱ�����������
#define ERR_TA_LOW_BALAN	-1212		//��������
#define ERR_TA_EXCEED_QUOTA	-1213		//���������޶�
#define ERR_TA_LOST_CARD	-1214		//��ʧ��
#define ERR_TA_FREEZE_CARD	-1215		//���Ῠ
#define ERR_TA_CARDNO		-1216		//�����ʺŲ���
#define ERR_TA_ID_CLOSE		-1217		//��ݹر�
#define ERR_TA_CR_DLL		-1218		//���ض�������̬���ӿ�ʧ��
#define ERR_TA_CR_INIT		-1219		//��������ʼ��ʧ��
#define ERR_TA_PARA			-1220		//��������
#define ERR_TA_NOREC		-1221		//û������ʻ�
#define ERR_TA_SUB_SUCC		-1222		//�����ɹ�,Ҳ����ȷ�ķ�����Ϣ
#define ERR_TA_SUB_FAIL		-1223		//����ʧ��,Ҳ����ȷ�ķ�����Ϣ
#define ERR_TA_INITED		-1224		//�������Ѿ���ʼ������ر�
#define ERR_TA_UP_DLL		-1225		//����UpdateAPI.dll��̬��ʧ��
#define ERR_TA_WRITECARD2	-1226		//��д�����ʧ��
#define ERR_TA_QUERYACC		-1227		//��ѯ�ʻ�ʧ��
#define ERR_TA_REREADCARD	-1228		//д��ʱ�����¶������ʧ��
#define ERR_TA_NO_FUN		-1229		//û���������

/*��������ͬ������������ֵ from -1300*/
#define ERR_SS_MAX_NG		-1300		//����ͬ������ʱ���������Ŀ
#define ERR_SS_NO_NG		-1301		//ɾ��ͬ������ʱû���ҵ�����
#define ERR_SS_DOWN_TSK		-1302		//����ͬ������ʱ����ͬ������ʧ��
#define ERR_SS_UP_TSK		-1303		//ɾ��ͬ������ʱ�ϴ�ͬ������ʧ��

#endif