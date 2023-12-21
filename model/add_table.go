package model

func addTable() {

	//1 alert_ignore
	tableAlertIgnore := AlertIgnore{}
	modelMap["alert_ignore"] = tableAlertIgnore
	primaryKeyMap["alert_ignore"] = tableAlertIgnore.PrimaryKey()
	//2 api_key_resource
	tableAPIKeyResource := APIKeyResource{}
	modelMap["api_key_resource"] = tableAPIKeyResource
	primaryKeyMap["api_key_resource"] = tableAPIKeyResource.PrimaryKey()
	//3 api_key_resource_privilege
	tableAPIKeyResourcePrivilege := APIKeyResourcePrivilege{}
	modelMap["api_key_resource_privilege"] = tableAPIKeyResourcePrivilege
	primaryKeyMap["api_key_resource_privilege"] = tableAPIKeyResourcePrivilege.PrimaryKey()
	//4 api_key_resource_type
	tableAPIKeyResourceType := APIKeyResourceType{}
	modelMap["api_key_resource_type"] = tableAPIKeyResourceType
	primaryKeyMap["api_key_resource_type"] = tableAPIKeyResourceType.PrimaryKey()
	//5 auth_group
	tableAuthGroup := AuthGroup{}
	modelMap["auth_group"] = tableAuthGroup
	primaryKeyMap["auth_group"] = tableAuthGroup.PrimaryKey()
	//6 auth_group_permissions
	tableAuthGroupPermissions := AuthGroupPermissions{}
	modelMap["auth_group_permissions"] = tableAuthGroupPermissions
	primaryKeyMap["auth_group_permissions"] = tableAuthGroupPermissions.PrimaryKey()
	//7 auth_permission
	tableAuthPermission := AuthPermission{}
	modelMap["auth_permission"] = tableAuthPermission
	primaryKeyMap["auth_permission"] = tableAuthPermission.PrimaryKey()
	//8 auth_user
	tableAuthUser := AuthUser{}
	modelMap["auth_user"] = tableAuthUser
	primaryKeyMap["auth_user"] = tableAuthUser.PrimaryKey()
	//9 auth_user_copy
	tableAuthUserCopy := AuthUserCopy{}
	modelMap["auth_user_copy"] = tableAuthUserCopy
	primaryKeyMap["auth_user_copy"] = tableAuthUserCopy.PrimaryKey()
	//10 auto_check_event
	tableAutoCheckEvent := AutoCheckEvent{}
	modelMap["auto_check_event"] = tableAutoCheckEvent
	primaryKeyMap["auto_check_event"] = tableAutoCheckEvent.PrimaryKey()
	//11 auto_check_host_info
	tableAutoCheckHostInfo := AutoCheckHostInfo{}
	modelMap["auto_check_host_info"] = tableAutoCheckHostInfo
	primaryKeyMap["auto_check_host_info"] = tableAutoCheckHostInfo.PrimaryKey()
	//12 auto_check_items
	tableAutoCheckItems := AutoCheckItems{}
	modelMap["auto_check_items"] = tableAutoCheckItems
	primaryKeyMap["auto_check_items"] = tableAutoCheckItems.PrimaryKey()
	//13 auto_check_items_detail
	tableAutoCheckItemsDetail := AutoCheckItemsDetail{}
	modelMap["auto_check_items_detail"] = tableAutoCheckItemsDetail
	primaryKeyMap["auto_check_items_detail"] = tableAutoCheckItemsDetail.PrimaryKey()
	//14 auto_check_result
	tableAutoCheckResult := AutoCheckResult{}
	modelMap["auto_check_result"] = tableAutoCheckResult
	primaryKeyMap["auto_check_result"] = tableAutoCheckResult.PrimaryKey()
	//15 auto_deploy
	tableAutoDeploy := AutoDeploy{}
	modelMap["auto_deploy"] = tableAutoDeploy
	primaryKeyMap["auto_deploy"] = tableAutoDeploy.PrimaryKey()
	//16 auto_test_record
	tableAutoTestRecord := AutoTestRecord{}
	modelMap["auto_test_record"] = tableAutoTestRecord
	primaryKeyMap["auto_test_record"] = tableAutoTestRecord.PrimaryKey()
	//17 backend_event_exe
	tableBackendEventExe := BackendEventExe{}
	modelMap["backend_event_exe"] = tableBackendEventExe
	primaryKeyMap["backend_event_exe"] = tableBackendEventExe.PrimaryKey()
	//18 bean_report
	tableBeanReport := BeanReport{}
	modelMap["bean_report"] = tableBeanReport
	primaryKeyMap["bean_report"] = tableBeanReport.PrimaryKey()
	//19 bean_report_issues
	tableBeanReportIssues := BeanReportIssues{}
	modelMap["bean_report_issues"] = tableBeanReportIssues
	primaryKeyMap["bean_report_issues"] = tableBeanReportIssues.PrimaryKey()
	//20 bean_report_team
	tableBeanReportTeam := BeanReportTeam{}
	modelMap["bean_report_team"] = tableBeanReportTeam
	primaryKeyMap["bean_report_team"] = tableBeanReportTeam.PrimaryKey()
	//21 black_list
	tableBlackList := BlackList{}
	modelMap["black_list"] = tableBlackList
	primaryKeyMap["black_list"] = tableBlackList.PrimaryKey()
	//22 build_flow
	tableBuildFlow := BuildFlow{}
	modelMap["build_flow"] = tableBuildFlow
	primaryKeyMap["build_flow"] = tableBuildFlow.PrimaryKey()
	//23 build_group
	tableBuildGroup := BuildGroup{}
	modelMap["build_group"] = tableBuildGroup
	primaryKeyMap["build_group"] = tableBuildGroup.PrimaryKey()
	//24 build_group_detail
	tableBuildGroupDetail := BuildGroupDetail{}
	modelMap["build_group_detail"] = tableBuildGroupDetail
	primaryKeyMap["build_group_detail"] = tableBuildGroupDetail.PrimaryKey()
	//25 chat_robot_instance
	tableChatRobotInstance := ChatRobotInstance{}
	modelMap["chat_robot_instance"] = tableChatRobotInstance
	primaryKeyMap["chat_robot_instance"] = tableChatRobotInstance.PrimaryKey()
	//26 collect_doc
	tableCollectDoc := CollectDoc{}
	modelMap["collect_doc"] = tableCollectDoc
	primaryKeyMap["collect_doc"] = tableCollectDoc.PrimaryKey()
	//27 collect_doc_demo
	tableCollectDocDemo := CollectDocDemo{}
	modelMap["collect_doc_demo"] = tableCollectDocDemo
	primaryKeyMap["collect_doc_demo"] = tableCollectDocDemo.PrimaryKey()
	//28 collect_doc_important
	tableCollectDocImportant := CollectDocImportant{}
	modelMap["collect_doc_important"] = tableCollectDocImportant
	primaryKeyMap["collect_doc_important"] = tableCollectDocImportant.PrimaryKey()
	//29 collect_doc_params
	tableCollectDocParams := CollectDocParams{}
	modelMap["collect_doc_params"] = tableCollectDocParams
	primaryKeyMap["collect_doc_params"] = tableCollectDocParams.PrimaryKey()
	//30 collect_doc_result
	tableCollectDocResult := CollectDocResult{}
	modelMap["collect_doc_result"] = tableCollectDocResult
	primaryKeyMap["collect_doc_result"] = tableCollectDocResult.PrimaryKey()
	//31 config_detail
	tableConfigDetail := ConfigDetail{}
	modelMap["config_detail"] = tableConfigDetail
	primaryKeyMap["config_detail"] = tableConfigDetail.PrimaryKey()
	//32 config_group
	tableConfigGroup := ConfigGroup{}
	modelMap["config_group"] = tableConfigGroup
	primaryKeyMap["config_group"] = tableConfigGroup.PrimaryKey()
	//33 dashboard_related_servers
	tableDashboardRelatedServers := DashboardRelatedServers{}
	modelMap["dashboard_related_servers"] = tableDashboardRelatedServers
	primaryKeyMap["dashboard_related_servers"] = tableDashboardRelatedServers.PrimaryKey()
	//34 dba_apk
	tableDbaApk := DbaApk{}
	modelMap["dba_apk"] = tableDbaApk
	primaryKeyMap["dba_apk"] = tableDbaApk.PrimaryKey()
	//35 dbdata_compare
	tableDbdataCompare := DbdataCompare{}
	modelMap["dbdata_compare"] = tableDbdataCompare
	primaryKeyMap["dbdata_compare"] = tableDbdataCompare.PrimaryKey()
	//36 dbdata_compare_change_list
	tableDbdataCompareChangeList := DbdataCompareChangeList{}
	modelMap["dbdata_compare_change_list"] = tableDbdataCompareChangeList
	primaryKeyMap["dbdata_compare_change_list"] = tableDbdataCompareChangeList.PrimaryKey()
	//37 dbdata_exp_table_standard
	tableDbdataExpTableStandard := DbdataExpTableStandard{}
	modelMap["dbdata_exp_table_standard"] = tableDbdataExpTableStandard
	primaryKeyMap["dbdata_exp_table_standard"] = tableDbdataExpTableStandard.PrimaryKey()
	//38 dbdata_export
	tableDbdataExport := DbdataExport{}
	modelMap["dbdata_export"] = tableDbdataExport
	primaryKeyMap["dbdata_export"] = tableDbdataExport.PrimaryKey()
	//39 dbdata_export_execute_log
	tableDbdataExportExecuteLog := DbdataExportExecuteLog{}
	modelMap["dbdata_export_execute_log"] = tableDbdataExportExecuteLog
	primaryKeyMap["dbdata_export_execute_log"] = tableDbdataExportExecuteLog.PrimaryKey()
	//40 dbdata_export_status
	tableDbdataExportStatus := DbdataExportStatus{}
	modelMap["dbdata_export_status"] = tableDbdataExportStatus
	primaryKeyMap["dbdata_export_status"] = tableDbdataExportStatus.PrimaryKey()
	//41 dbdata_export_tables
	tableDbdataExportTables := DbdataExportTables{}
	modelMap["dbdata_export_tables"] = tableDbdataExportTables
	primaryKeyMap["dbdata_export_tables"] = tableDbdataExportTables.PrimaryKey()
	//42 dbdata_export_users
	tableDbdataExportUsers := DbdataExportUsers{}
	modelMap["dbdata_export_users"] = tableDbdataExportUsers
	primaryKeyMap["dbdata_export_users"] = tableDbdataExportUsers.PrimaryKey()
	//43 dbdata_import
	tableDbdataImport := DbdataImport{}
	modelMap["dbdata_import"] = tableDbdataImport
	primaryKeyMap["dbdata_import"] = tableDbdataImport.PrimaryKey()
	//44 dbdata_import_execute_log
	tableDbdataImportExecuteLog := DbdataImportExecuteLog{}
	modelMap["dbdata_import_execute_log"] = tableDbdataImportExecuteLog
	primaryKeyMap["dbdata_import_execute_log"] = tableDbdataImportExecuteLog.PrimaryKey()
	//45 dbdata_import_status
	tableDbdataImportStatus := DbdataImportStatus{}
	modelMap["dbdata_import_status"] = tableDbdataImportStatus
	primaryKeyMap["dbdata_import_status"] = tableDbdataImportStatus.PrimaryKey()
	//46 dbdata_import_tables
	tableDbdataImportTables := DbdataImportTables{}
	modelMap["dbdata_import_tables"] = tableDbdataImportTables
	primaryKeyMap["dbdata_import_tables"] = tableDbdataImportTables.PrimaryKey()
	//47 dbdata_import_users
	tableDbdataImportUsers := DbdataImportUsers{}
	modelMap["dbdata_import_users"] = tableDbdataImportUsers
	primaryKeyMap["dbdata_import_users"] = tableDbdataImportUsers.PrimaryKey()
	//48 dbdata_sync_basetable
	tableDbdataSyncBasetable := DbdataSyncBasetable{}
	modelMap["dbdata_sync_basetable"] = tableDbdataSyncBasetable
	primaryKeyMap["dbdata_sync_basetable"] = tableDbdataSyncBasetable.PrimaryKey()
	//49 dbdata_sync_exp_event
	tableDbdataSyncExpEvent := DbdataSyncExpEvent{}
	modelMap["dbdata_sync_exp_event"] = tableDbdataSyncExpEvent
	primaryKeyMap["dbdata_sync_exp_event"] = tableDbdataSyncExpEvent.PrimaryKey()
	//50 dbdata_sync_imp_event
	tableDbdataSyncImpEvent := DbdataSyncImpEvent{}
	modelMap["dbdata_sync_imp_event"] = tableDbdataSyncImpEvent
	primaryKeyMap["dbdata_sync_imp_event"] = tableDbdataSyncImpEvent.PrimaryKey()
	//51 dbdata_sync_step_event
	tableDbdataSyncStepEvent := DbdataSyncStepEvent{}
	modelMap["dbdata_sync_step_event"] = tableDbdataSyncStepEvent
	primaryKeyMap["dbdata_sync_step_event"] = tableDbdataSyncStepEvent.PrimaryKey()
	//52 dbscript_issue_list
	tableDbscriptIssueList := DbscriptIssueList{}
	modelMap["dbscript_issue_list"] = tableDbscriptIssueList
	primaryKeyMap["dbscript_issue_list"] = tableDbscriptIssueList.PrimaryKey()
	//53 dbscript_req
	tableDbscriptReq := DbscriptReq{}
	modelMap["dbscript_req"] = tableDbscriptReq
	primaryKeyMap["dbscript_req"] = tableDbscriptReq.PrimaryKey()
	//54 dbscript_req_event
	tableDbscriptReqEvent := DbscriptReqEvent{}
	modelMap["dbscript_req_event"] = tableDbscriptReqEvent
	primaryKeyMap["dbscript_req_event"] = tableDbscriptReqEvent.PrimaryKey()
	//55 dbscript_req_info
	tableDbscriptReqInfo := DbscriptReqInfo{}
	modelMap["dbscript_req_info"] = tableDbscriptReqInfo
	primaryKeyMap["dbscript_req_info"] = tableDbscriptReqInfo.PrimaryKey()
	//56 dbscript_req_project
	tableDbscriptReqProject := DbscriptReqProject{}
	modelMap["dbscript_req_project"] = tableDbscriptReqProject
	primaryKeyMap["dbscript_req_project"] = tableDbscriptReqProject.PrimaryKey()
	//57 dep_local_link
	tableDepLocalLink := DepLocalLink{}
	modelMap["dep_local_link"] = tableDepLocalLink
	primaryKeyMap["dep_local_link"] = tableDepLocalLink.PrimaryKey()
	//58 dep_task_update_record
	tableDepTaskUpdateRecord := DepTaskUpdateRecord{}
	modelMap["dep_task_update_record"] = tableDepTaskUpdateRecord
	primaryKeyMap["dep_task_update_record"] = tableDepTaskUpdateRecord.PrimaryKey()
	//59 dep_tasklist_update_record
	tableDepTasklistUpdateRecord := DepTasklistUpdateRecord{}
	modelMap["dep_tasklist_update_record"] = tableDepTasklistUpdateRecord
	primaryKeyMap["dep_tasklist_update_record"] = tableDepTasklistUpdateRecord.PrimaryKey()
	//60 deploy_conf_diff_detail
	tableDeployConfDiffDetail := DeployConfDiffDetail{}
	modelMap["deploy_conf_diff_detail"] = tableDeployConfDiffDetail
	primaryKeyMap["deploy_conf_diff_detail"] = tableDeployConfDiffDetail.PrimaryKey()
	//61 deploy_env_conf_template
	tableDeployEnvConfTemplate := DeployEnvConfTemplate{}
	modelMap["deploy_env_conf_template"] = tableDeployEnvConfTemplate
	primaryKeyMap["deploy_env_conf_template"] = tableDeployEnvConfTemplate.PrimaryKey()
	//62 deploy_event
	tableDeployEvent := DeployEvent{}
	modelMap["deploy_event"] = tableDeployEvent
	primaryKeyMap["deploy_event"] = tableDeployEvent.PrimaryKey()
	//63 deploy_favorite_war
	tableDeployFavoriteWar := DeployFavoriteWar{}
	modelMap["deploy_favorite_war"] = tableDeployFavoriteWar
	primaryKeyMap["deploy_favorite_war"] = tableDeployFavoriteWar.PrimaryKey()
	//64 deploy_prepare_event
	tableDeployPrepareEvent := DeployPrepareEvent{}
	modelMap["deploy_prepare_event"] = tableDeployPrepareEvent
	primaryKeyMap["deploy_prepare_event"] = tableDeployPrepareEvent.PrimaryKey()
	//65 deploy_replace_event
	tableDeployReplaceEvent := DeployReplaceEvent{}
	modelMap["deploy_replace_event"] = tableDeployReplaceEvent
	primaryKeyMap["deploy_replace_event"] = tableDeployReplaceEvent.PrimaryKey()
	//66 deploy_replace_war
	tableDeployReplaceWar := DeployReplaceWar{}
	modelMap["deploy_replace_war"] = tableDeployReplaceWar
	primaryKeyMap["deploy_replace_war"] = tableDeployReplaceWar.PrimaryKey()
	//67 deploy_req_doc
	tableDeployReqDoc := DeployReqDoc{}
	modelMap["deploy_req_doc"] = tableDeployReqDoc
	primaryKeyMap["deploy_req_doc"] = tableDeployReqDoc.PrimaryKey()
	//68 deploy_task
	tableDeployTask := DeployTask{}
	modelMap["deploy_task"] = tableDeployTask
	primaryKeyMap["deploy_task"] = tableDeployTask.PrimaryKey()
	//69 deploy_task_once
	tableDeployTaskOnce := DeployTaskOnce{}
	modelMap["deploy_task_once"] = tableDeployTaskOnce
	primaryKeyMap["deploy_task_once"] = tableDeployTaskOnce.PrimaryKey()
	//70 deploy_task_once_flow
	tableDeployTaskOnceFlow := DeployTaskOnceFlow{}
	modelMap["deploy_task_once_flow"] = tableDeployTaskOnceFlow
	primaryKeyMap["deploy_task_once_flow"] = tableDeployTaskOnceFlow.PrimaryKey()
	//71 deploy_tasklist
	tableDeployTasklist := DeployTasklist{}
	modelMap["deploy_tasklist"] = tableDeployTasklist
	primaryKeyMap["deploy_tasklist"] = tableDeployTasklist.PrimaryKey()
	//72 deptask_springboot_conf
	tableDeptaskSpringbootConf := DeptaskSpringbootConf{}
	modelMap["deptask_springboot_conf"] = tableDeptaskSpringbootConf
	primaryKeyMap["deptask_springboot_conf"] = tableDeptaskSpringbootConf.PrimaryKey()
	//73 dish_bind
	tableDishBind := DishBind{}
	modelMap["dish_bind"] = tableDishBind
	primaryKeyMap["dish_bind"] = tableDishBind.PrimaryKey()
	//74 dish_type
	tableDishType := DishType{}
	modelMap["dish_type"] = tableDishType
	primaryKeyMap["dish_type"] = tableDishType.PrimaryKey()
	//75 dish_type_rule
	tableDishTypeRule := DishTypeRule{}
	modelMap["dish_type_rule"] = tableDishTypeRule
	primaryKeyMap["dish_type_rule"] = tableDishTypeRule.PrimaryKey()
	//76 django_admin_log
	tableDjangoAdminLog := DjangoAdminLog{}
	modelMap["django_admin_log"] = tableDjangoAdminLog
	primaryKeyMap["django_admin_log"] = tableDjangoAdminLog.PrimaryKey()
	//77 django_apscheduler_djangojob
	tableDjangoApschedulerDjangojob := DjangoApschedulerDjangojob{}
	modelMap["django_apscheduler_djangojob"] = tableDjangoApschedulerDjangojob
	primaryKeyMap["django_apscheduler_djangojob"] = tableDjangoApschedulerDjangojob.PrimaryKey()
	//78 django_apscheduler_djangojobexecution
	tableDjangoApschedulerDjangojobexecution := DjangoApschedulerDjangojobexecution{}
	modelMap["django_apscheduler_djangojobexecution"] = tableDjangoApschedulerDjangojobexecution
	primaryKeyMap["django_apscheduler_djangojobexecution"] = tableDjangoApschedulerDjangojobexecution.PrimaryKey()
	//79 django_content_type
	tableDjangoContentType := DjangoContentType{}
	modelMap["django_content_type"] = tableDjangoContentType
	primaryKeyMap["django_content_type"] = tableDjangoContentType.PrimaryKey()
	//80 django_migrations
	tableDjangoMigrations := DjangoMigrations{}
	modelMap["django_migrations"] = tableDjangoMigrations
	primaryKeyMap["django_migrations"] = tableDjangoMigrations.PrimaryKey()
	//81 django_session
	tableDjangoSession := DjangoSession{}
	modelMap["django_session"] = tableDjangoSession
	primaryKeyMap["django_session"] = tableDjangoSession.PrimaryKey()
	//82 djiango_cache
	tableDjiangoCache := DjiangoCache{}
	modelMap["djiango_cache"] = tableDjiangoCache
	primaryKeyMap["djiango_cache"] = tableDjiangoCache.PrimaryKey()
	//83 doc_group
	tableDocGroup := DocGroup{}
	modelMap["doc_group"] = tableDocGroup
	primaryKeyMap["doc_group"] = tableDocGroup.PrimaryKey()
	//84 dump_history
	tableDumpHistory := DumpHistory{}
	modelMap["dump_history"] = tableDumpHistory
	primaryKeyMap["dump_history"] = tableDumpHistory.PrimaryKey()
	//85 emergency_history
	tableEmergencyHistory := EmergencyHistory{}
	modelMap["emergency_history"] = tableEmergencyHistory
	primaryKeyMap["emergency_history"] = tableEmergencyHistory.PrimaryKey()
	//86 ens_activity_data_days
	tableEnsActivityDataDays := EnsActivityDataDays{}
	modelMap["ens_activity_data_days"] = tableEnsActivityDataDays
	primaryKeyMap["ens_activity_data_days"] = tableEnsActivityDataDays.PrimaryKey()
	//87 ens_activity_data_hours
	tableEnsActivityDataHours := EnsActivityDataHours{}
	modelMap["ens_activity_data_hours"] = tableEnsActivityDataHours
	primaryKeyMap["ens_activity_data_hours"] = tableEnsActivityDataHours.PrimaryKey()
	//88 ens_activity_data_seconds
	tableEnsActivityDataSeconds := EnsActivityDataSeconds{}
	modelMap["ens_activity_data_seconds"] = tableEnsActivityDataSeconds
	primaryKeyMap["ens_activity_data_seconds"] = tableEnsActivityDataSeconds.PrimaryKey()
	//89 ens_channelinfos
	tableEnsChannelinfos := EnsChannelinfos{}
	modelMap["ens_channelinfos"] = tableEnsChannelinfos
	primaryKeyMap["ens_channelinfos"] = tableEnsChannelinfos.PrimaryKey()
	//90 ens_messageheader
	tableEnsMessageheader := EnsMessageheader{}
	modelMap["ens_messageheader"] = tableEnsMessageheader
	primaryKeyMap["ens_messageheader"] = tableEnsMessageheader.PrimaryKey()
	//91 ens_messageheader_hour
	tableEnsMessageheaderHour := EnsMessageheaderHour{}
	modelMap["ens_messageheader_hour"] = tableEnsMessageheaderHour
	primaryKeyMap["ens_messageheader_hour"] = tableEnsMessageheaderHour.PrimaryKey()
	//92 ens_util_log
	tableEnsUtilLog := EnsUtilLog{}
	modelMap["ens_util_log"] = tableEnsUtilLog
	primaryKeyMap["ens_util_log"] = tableEnsUtilLog.PrimaryKey()
	//93 env_variable
	tableEnvVariable := EnvVariable{}
	modelMap["env_variable"] = tableEnvVariable
	primaryKeyMap["env_variable"] = tableEnvVariable.PrimaryKey()
	//94 event_log
	tableEventLog := EventLog{}
	modelMap["event_log"] = tableEventLog
	primaryKeyMap["event_log"] = tableEventLog.PrimaryKey()
	//95 explain_issues
	tableExplainIssues := ExplainIssues{}
	modelMap["explain_issues"] = tableExplainIssues
	primaryKeyMap["explain_issues"] = tableExplainIssues.PrimaryKey()
	//96 explain_planning
	tableExplainPlanning := ExplainPlanning{}
	modelMap["explain_planning"] = tableExplainPlanning
	primaryKeyMap["explain_planning"] = tableExplainPlanning.PrimaryKey()
	//97 explain_users
	tableExplainUsers := ExplainUsers{}
	modelMap["explain_users"] = tableExplainUsers
	primaryKeyMap["explain_users"] = tableExplainUsers.PrimaryKey()
	//98 failure_report
	tableFailureReport := FailureReport{}
	modelMap["failure_report"] = tableFailureReport
	primaryKeyMap["failure_report"] = tableFailureReport.PrimaryKey()
	//99 failure_report_change_list
	tableFailureReportChangeList := FailureReportChangeList{}
	modelMap["failure_report_change_list"] = tableFailureReportChangeList
	primaryKeyMap["failure_report_change_list"] = tableFailureReportChangeList.PrimaryKey()
	//100 failure_report_second
	tableFailureReportSecond := FailureReportSecond{}
	modelMap["failure_report_second"] = tableFailureReportSecond
	primaryKeyMap["failure_report_second"] = tableFailureReportSecond.PrimaryKey()
	//101 file_data
	tableFileData := FileData{}
	modelMap["file_data"] = tableFileData
	primaryKeyMap["file_data"] = tableFileData.PrimaryKey()
	//102 fr_errrecord
	tableFrErrrecord := FrErrrecord{}
	modelMap["fr_errrecord"] = tableFrErrrecord
	primaryKeyMap["fr_errrecord"] = tableFrErrrecord.PrimaryKey()
	//103 fr_exerecord
	tableFrExerecord := FrExerecord{}
	modelMap["fr_exerecord"] = tableFrExerecord
	primaryKeyMap["fr_exerecord"] = tableFrExerecord.PrimaryKey()
	//104 fr_main_core
	tableFrMainCore := FrMainCore{}
	modelMap["fr_main_core"] = tableFrMainCore
	primaryKeyMap["fr_main_core"] = tableFrMainCore.PrimaryKey()
	//105 git_pdm_info
	tableGitPdmInfo := GitPdmInfo{}
	modelMap["git_pdm_info"] = tableGitPdmInfo
	primaryKeyMap["git_pdm_info"] = tableGitPdmInfo.PrimaryKey()
	//106 global_conf_details
	tableGlobalConfDetails := GlobalConfDetails{}
	modelMap["global_conf_details"] = tableGlobalConfDetails
	primaryKeyMap["global_conf_details"] = tableGlobalConfDetails.PrimaryKey()
	//107 global_conf_group
	tableGlobalConfGroup := GlobalConfGroup{}
	modelMap["global_conf_group"] = tableGlobalConfGroup
	primaryKeyMap["global_conf_group"] = tableGlobalConfGroup.PrimaryKey()
	//108 global_conf_modify_log
	tableGlobalConfModifyLog := GlobalConfModifyLog{}
	modelMap["global_conf_modify_log"] = tableGlobalConfModifyLog
	primaryKeyMap["global_conf_modify_log"] = tableGlobalConfModifyLog.PrimaryKey()
	//109 global_item_explain
	tableGlobalItemExplain := GlobalItemExplain{}
	modelMap["global_item_explain"] = tableGlobalItemExplain
	primaryKeyMap["global_item_explain"] = tableGlobalItemExplain.PrimaryKey()
	//110 his_issue_record
	tableHisIssueRecord := HisIssueRecord{}
	modelMap["his_issue_record"] = tableHisIssueRecord
	primaryKeyMap["his_issue_record"] = tableHisIssueRecord.PrimaryKey()
	//111 his_issue_record_img
	tableHisIssueRecordImg := HisIssueRecordImg{}
	modelMap["his_issue_record_img"] = tableHisIssueRecordImg
	primaryKeyMap["his_issue_record_img"] = tableHisIssueRecordImg.PrimaryKey()
	//112 his_issue_statu
	tableHisIssueStatu := HisIssueStatu{}
	modelMap["his_issue_statu"] = tableHisIssueStatu
	primaryKeyMap["his_issue_statu"] = tableHisIssueStatu.PrimaryKey()
	//113 his_status_change_history
	tableHisStatusChangeHistory := HisStatusChangeHistory{}
	modelMap["his_status_change_history"] = tableHisStatusChangeHistory
	primaryKeyMap["his_status_change_history"] = tableHisStatusChangeHistory.PrimaryKey()
	//114 icon_data
	tableIconData := IconData{}
	modelMap["icon_data"] = tableIconData
	primaryKeyMap["icon_data"] = tableIconData.PrimaryKey()
	//115 icon_group
	tableIconGroup := IconGroup{}
	modelMap["icon_group"] = tableIconGroup
	primaryKeyMap["icon_group"] = tableIconGroup.PrimaryKey()
	//116 icon_project
	tableIconProject := IconProject{}
	modelMap["icon_project"] = tableIconProject
	primaryKeyMap["icon_project"] = tableIconProject.PrimaryKey()
	//117 icon_tag
	tableIconTag := IconTag{}
	modelMap["icon_tag"] = tableIconTag
	primaryKeyMap["icon_tag"] = tableIconTag.PrimaryKey()
	//118 install_os_user_801
	tableInstallOsUser801 := InstallOsUser801{}
	modelMap["install_os_user_801"] = tableInstallOsUser801
	primaryKeyMap["install_os_user_801"] = tableInstallOsUser801.PrimaryKey()
	//119 ip_identify_middle
	tableIPIdentifyMiddle := IPIdentifyMiddle{}
	modelMap["ip_identify_middle"] = tableIPIdentifyMiddle
	primaryKeyMap["ip_identify_middle"] = tableIPIdentifyMiddle.PrimaryKey()
	//120 issue_commit
	tableIssueCommit := IssueCommit{}
	modelMap["issue_commit"] = tableIssueCommit
	primaryKeyMap["issue_commit"] = tableIssueCommit.PrimaryKey()
	//121 issue_flow
	tableIssueFlow := IssueFlow{}
	modelMap["issue_flow"] = tableIssueFlow
	primaryKeyMap["issue_flow"] = tableIssueFlow.PrimaryKey()
	//122 issue_fow
	tableIssueFow := IssueFow{}
	modelMap["issue_fow"] = tableIssueFow
	primaryKeyMap["issue_fow"] = tableIssueFow.PrimaryKey()
	//123 issue_img
	tableIssueImg := IssueImg{}
	modelMap["issue_img"] = tableIssueImg
	primaryKeyMap["issue_img"] = tableIssueImg.PrimaryKey()
	//124 issue_record_img
	tableIssueRecordImg := IssueRecordImg{}
	modelMap["issue_record_img"] = tableIssueRecordImg
	primaryKeyMap["issue_record_img"] = tableIssueRecordImg.PrimaryKey()
	//125 it_resource_node
	tableItResourceNode := ItResourceNode{}
	modelMap["it_resource_node"] = tableItResourceNode
	primaryKeyMap["it_resource_node"] = tableItResourceNode.PrimaryKey()
	//126 items
	tableItems := Items{}
	modelMap["items"] = tableItems
	primaryKeyMap["items"] = tableItems.PrimaryKey()
	//127 jbosslogcategory_operator
	tableJbosslogcategoryOperator := JbosslogcategoryOperator{}
	modelMap["jbosslogcategory_operator"] = tableJbosslogcategoryOperator
	primaryKeyMap["jbosslogcategory_operator"] = tableJbosslogcategoryOperator.PrimaryKey()
	//128 jira_backend_status
	tableJiraBackendStatus := JiraBackendStatus{}
	modelMap["jira_backend_status"] = tableJiraBackendStatus
	primaryKeyMap["jira_backend_status"] = tableJiraBackendStatus.PrimaryKey()
	//129 jira_issue_slice
	tableJiraIssueSlice := JiraIssueSlice{}
	modelMap["jira_issue_slice"] = tableJiraIssueSlice
	primaryKeyMap["jira_issue_slice"] = tableJiraIssueSlice.PrimaryKey()
	//130 jira_op_log
	tableJiraOpLog := JiraOpLog{}
	modelMap["jira_op_log"] = tableJiraOpLog
	primaryKeyMap["jira_op_log"] = tableJiraOpLog.PrimaryKey()
	//131 jobs_exec_log
	tableJobsExecLog := JobsExecLog{}
	modelMap["jobs_exec_log"] = tableJobsExecLog
	primaryKeyMap["jobs_exec_log"] = tableJobsExecLog.PrimaryKey()
	//132 jobs_list
	tableJobsList := JobsList{}
	modelMap["jobs_list"] = tableJobsList
	primaryKeyMap["jobs_list"] = tableJobsList.PrimaryKey()
	//133 jobs_rel_server
	tableJobsRelServer := JobsRelServer{}
	modelMap["jobs_rel_server"] = tableJobsRelServer
	primaryKeyMap["jobs_rel_server"] = tableJobsRelServer.PrimaryKey()
	//134 ldap_group
	tableLdapGroup := LdapGroup{}
	modelMap["ldap_group"] = tableLdapGroup
	primaryKeyMap["ldap_group"] = tableLdapGroup.PrimaryKey()
	//135 m_test1
	tableMTest1 := MTest1{}
	modelMap["m_test1"] = tableMTest1
	primaryKeyMap["m_test1"] = tableMTest1.PrimaryKey()
	//136 merge_publish
	tableMergePublish := MergePublish{}
	modelMap["merge_publish"] = tableMergePublish
	primaryKeyMap["merge_publish"] = tableMergePublish.PrimaryKey()
	//137 merge_publish_detail
	tableMergePublishDetail := MergePublishDetail{}
	modelMap["merge_publish_detail"] = tableMergePublishDetail
	primaryKeyMap["merge_publish_detail"] = tableMergePublishDetail.PrimaryKey()
	//138 merge_publish_war_detail
	tableMergePublishWarDetail := MergePublishWarDetail{}
	modelMap["merge_publish_war_detail"] = tableMergePublishWarDetail
	primaryKeyMap["merge_publish_war_detail"] = tableMergePublishWarDetail.PrimaryKey()
	//139 merge_request_history
	tableMergeRequestHistory := MergeRequestHistory{}
	modelMap["merge_request_history"] = tableMergeRequestHistory
	primaryKeyMap["merge_request_history"] = tableMergeRequestHistory.PrimaryKey()
	//140 message_send_log
	tableMessageSendLog := MessageSendLog{}
	modelMap["message_send_log"] = tableMessageSendLog
	primaryKeyMap["message_send_log"] = tableMessageSendLog.PrimaryKey()
	//141 message_send_serverEnv
	tableMessageSendServerEnv := MessageSendServerEnv{}
	modelMap["message_send_serverEnv"] = tableMessageSendServerEnv
	primaryKeyMap["message_send_serverEnv"] = tableMessageSendServerEnv.PrimaryKey()
	//142 message_send_target
	tableMessageSendTarget := MessageSendTarget{}
	modelMap["message_send_target"] = tableMessageSendTarget
	primaryKeyMap["message_send_target"] = tableMessageSendTarget.PrimaryKey()
	//143 monitor_active_test
	tableMonitorActiveTest := MonitorActiveTest{}
	modelMap["monitor_active_test"] = tableMonitorActiveTest
	primaryKeyMap["monitor_active_test"] = tableMonitorActiveTest.PrimaryKey()
	//144 monitor_agent
	tableMonitorAgent := MonitorAgent{}
	modelMap["monitor_agent"] = tableMonitorAgent
	primaryKeyMap["monitor_agent"] = tableMonitorAgent.PrimaryKey()
	//145 monitor_alertitem_confirm
	tableMonitorAlertitemConfirm := MonitorAlertitemConfirm{}
	modelMap["monitor_alertitem_confirm"] = tableMonitorAlertitemConfirm
	primaryKeyMap["monitor_alertitem_confirm"] = tableMonitorAlertitemConfirm.PrimaryKey()
	//146 monitor_alertitem_record
	tableMonitorAlertitemRecord := MonitorAlertitemRecord{}
	modelMap["monitor_alertitem_record"] = tableMonitorAlertitemRecord
	primaryKeyMap["monitor_alertitem_record"] = tableMonitorAlertitemRecord.PrimaryKey()
	//147 monitor_alertitem_solution
	tableMonitorAlertitemSolution := MonitorAlertitemSolution{}
	modelMap["monitor_alertitem_solution"] = tableMonitorAlertitemSolution
	primaryKeyMap["monitor_alertitem_solution"] = tableMonitorAlertitemSolution.PrimaryKey()
	//148 monitor_item_record_data
	tableMonitorItemRecordData := MonitorItemRecordData{}
	modelMap["monitor_item_record_data"] = tableMonitorItemRecordData
	primaryKeyMap["monitor_item_record_data"] = tableMonitorItemRecordData.PrimaryKey()
	//149 monitor_itemkey_desc
	tableMonitorItemkeyDesc := MonitorItemkeyDesc{}
	modelMap["monitor_itemkey_desc"] = tableMonitorItemkeyDesc
	primaryKeyMap["monitor_itemkey_desc"] = tableMonitorItemkeyDesc.PrimaryKey()
	//150 monitor_msg_channel
	tableMonitorMsgChannel := MonitorMsgChannel{}
	modelMap["monitor_msg_channel"] = tableMonitorMsgChannel
	primaryKeyMap["monitor_msg_channel"] = tableMonitorMsgChannel.PrimaryKey()
	//151 monitor_msg_channel_group
	tableMonitorMsgChannelGroup := MonitorMsgChannelGroup{}
	modelMap["monitor_msg_channel_group"] = tableMonitorMsgChannelGroup
	primaryKeyMap["monitor_msg_channel_group"] = tableMonitorMsgChannelGroup.PrimaryKey()
	//152 monitor_msg_channel_group_rel
	tableMonitorMsgChannelGroupRel := MonitorMsgChannelGroupRel{}
	modelMap["monitor_msg_channel_group_rel"] = tableMonitorMsgChannelGroupRel
	primaryKeyMap["monitor_msg_channel_group_rel"] = tableMonitorMsgChannelGroupRel.PrimaryKey()
	//153 monitor_name_code
	tableMonitorNameCode := MonitorNameCode{}
	modelMap["monitor_name_code"] = tableMonitorNameCode
	primaryKeyMap["monitor_name_code"] = tableMonitorNameCode.PrimaryKey()
	//154 monitor_recovery_access
	tableMonitorRecoveryAccess := MonitorRecoveryAccess{}
	modelMap["monitor_recovery_access"] = tableMonitorRecoveryAccess
	primaryKeyMap["monitor_recovery_access"] = tableMonitorRecoveryAccess.PrimaryKey()
	//155 monitor_recovery_access_rel
	tableMonitorRecoveryAccessRel := MonitorRecoveryAccessRel{}
	modelMap["monitor_recovery_access_rel"] = tableMonitorRecoveryAccessRel
	primaryKeyMap["monitor_recovery_access_rel"] = tableMonitorRecoveryAccessRel.PrimaryKey()
	//156 monitor_server_dashboard
	tableMonitorServerDashboard := MonitorServerDashboard{}
	modelMap["monitor_server_dashboard"] = tableMonitorServerDashboard
	primaryKeyMap["monitor_server_dashboard"] = tableMonitorServerDashboard.PrimaryKey()
	//157 moongod_event_exe
	tableMoongodEventExe := MoongodEventExe{}
	modelMap["moongod_event_exe"] = tableMoongodEventExe
	primaryKeyMap["moongod_event_exe"] = tableMoongodEventExe.PrimaryKey()
	//158 mr_token
	tableMrToken := MrToken{}
	modelMap["mr_token"] = tableMrToken
	primaryKeyMap["mr_token"] = tableMrToken.PrimaryKey()
	//159 msg_send_forbid
	tableMsgSendForbid := MsgSendForbid{}
	modelMap["msg_send_forbid"] = tableMsgSendForbid
	primaryKeyMap["msg_send_forbid"] = tableMsgSendForbid.PrimaryKey()
	//160 msg_send_log
	tableMsgSendLog := MsgSendLog{}
	modelMap["msg_send_log"] = tableMsgSendLog
	primaryKeyMap["msg_send_log"] = tableMsgSendLog.PrimaryKey()
	//161 msg_send_rel_channel
	tableMsgSendRelChannel := MsgSendRelChannel{}
	modelMap["msg_send_rel_channel"] = tableMsgSendRelChannel
	primaryKeyMap["msg_send_rel_channel"] = tableMsgSendRelChannel.PrimaryKey()
	//162 msg_send_rule
	tableMsgSendRule := MsgSendRule{}
	modelMap["msg_send_rule"] = tableMsgSendRule
	primaryKeyMap["msg_send_rule"] = tableMsgSendRule.PrimaryKey()
	//163 msg_send_rule_group
	tableMsgSendRuleGroup := MsgSendRuleGroup{}
	modelMap["msg_send_rule_group"] = tableMsgSendRuleGroup
	primaryKeyMap["msg_send_rule_group"] = tableMsgSendRuleGroup.PrimaryKey()
	//164 msg_send_token
	tableMsgSendToken := MsgSendToken{}
	modelMap["msg_send_token"] = tableMsgSendToken
	primaryKeyMap["msg_send_token"] = tableMsgSendToken.PrimaryKey()
	//165 msg_sendchannel_provider
	tableMsgSendchannelProvider := MsgSendchannelProvider{}
	modelMap["msg_sendchannel_provider"] = tableMsgSendchannelProvider
	primaryKeyMap["msg_sendchannel_provider"] = tableMsgSendchannelProvider.PrimaryKey()
	//166 my_focus_dashboard
	tableMyFocusDashboard := MyFocusDashboard{}
	modelMap["my_focus_dashboard"] = tableMyFocusDashboard
	primaryKeyMap["my_focus_dashboard"] = tableMyFocusDashboard.PrimaryKey()
	//167 notice_server_env
	tableNoticeServerEnv := NoticeServerEnv{}
	modelMap["notice_server_env"] = tableNoticeServerEnv
	primaryKeyMap["notice_server_env"] = tableNoticeServerEnv.PrimaryKey()
	//168 open_api
	tableOpenAPI := OpenAPI{}
	modelMap["open_api"] = tableOpenAPI
	primaryKeyMap["open_api"] = tableOpenAPI.PrimaryKey()
	//169 open_api_key
	tableOpenAPIKey := OpenAPIKey{}
	modelMap["open_api_key"] = tableOpenAPIKey
	primaryKeyMap["open_api_key"] = tableOpenAPIKey.PrimaryKey()
	//170 open_api_log
	tableOpenAPILog := OpenAPILog{}
	modelMap["open_api_log"] = tableOpenAPILog
	primaryKeyMap["open_api_log"] = tableOpenAPILog.PrimaryKey()
	//171 open_job_event
	tableOpenJobEvent := OpenJobEvent{}
	modelMap["open_job_event"] = tableOpenJobEvent
	primaryKeyMap["open_job_event"] = tableOpenJobEvent.PrimaryKey()
	//172 open_job_list
	tableOpenJobList := OpenJobList{}
	modelMap["open_job_list"] = tableOpenJobList
	primaryKeyMap["open_job_list"] = tableOpenJobList.PrimaryKey()
	//173 ops_config_detail
	tableOpsConfigDetail := OpsConfigDetail{}
	modelMap["ops_config_detail"] = tableOpsConfigDetail
	primaryKeyMap["ops_config_detail"] = tableOpsConfigDetail.PrimaryKey()
	//174 ops_config_detail_copy
	tableOpsConfigDetailCopy := OpsConfigDetailCopy{}
	modelMap["ops_config_detail_copy"] = tableOpsConfigDetailCopy
	primaryKeyMap["ops_config_detail_copy"] = tableOpsConfigDetailCopy.PrimaryKey()
	//175 ops_config_group
	tableOpsConfigGroup := OpsConfigGroup{}
	modelMap["ops_config_group"] = tableOpsConfigGroup
	primaryKeyMap["ops_config_group"] = tableOpsConfigGroup.PrimaryKey()
	//176 ops_software_list
	tableOpsSoftwareList := OpsSoftwareList{}
	modelMap["ops_software_list"] = tableOpsSoftwareList
	primaryKeyMap["ops_software_list"] = tableOpsSoftwareList.PrimaryKey()
	//177 ops_sw_deploy_event
	tableOpsSwDeployEvent := OpsSwDeployEvent{}
	modelMap["ops_sw_deploy_event"] = tableOpsSwDeployEvent
	primaryKeyMap["ops_sw_deploy_event"] = tableOpsSwDeployEvent.PrimaryKey()
	//178 ops_sw_deploy_list
	tableOpsSwDeployList := OpsSwDeployList{}
	modelMap["ops_sw_deploy_list"] = tableOpsSwDeployList
	primaryKeyMap["ops_sw_deploy_list"] = tableOpsSwDeployList.PrimaryKey()
	//179 oracle_awr
	tableOracleAwr := OracleAwr{}
	modelMap["oracle_awr"] = tableOracleAwr
	primaryKeyMap["oracle_awr"] = tableOracleAwr.PrimaryKey()
	//180 oracle_instance
	tableOracleInstance := OracleInstance{}
	modelMap["oracle_instance"] = tableOracleInstance
	primaryKeyMap["oracle_instance"] = tableOracleInstance.PrimaryKey()
	//181 oracle_rpt
	tableOracleRpt := OracleRpt{}
	modelMap["oracle_rpt"] = tableOracleRpt
	primaryKeyMap["oracle_rpt"] = tableOracleRpt.PrimaryKey()
	//182 oracle_user
	tableOracleUser := OracleUser{}
	modelMap["oracle_user"] = tableOracleUser
	primaryKeyMap["oracle_user"] = tableOracleUser.PrimaryKey()
	//183 os_soft_user_group
	tableOsSoftUserGroup := OsSoftUserGroup{}
	modelMap["os_soft_user_group"] = tableOsSoftUserGroup
	primaryKeyMap["os_soft_user_group"] = tableOsSoftUserGroup.PrimaryKey()
	//184 page_group
	tablePageGroup := PageGroup{}
	modelMap["page_group"] = tablePageGroup
	primaryKeyMap["page_group"] = tablePageGroup.PrimaryKey()
	//185 page_group_item
	tablePageGroupItem := PageGroupItem{}
	modelMap["page_group_item"] = tablePageGroupItem
	primaryKeyMap["page_group_item"] = tablePageGroupItem.PrimaryKey()
	//186 pinpoint_app_name
	tablePinpointAppName := PinpointAppName{}
	modelMap["pinpoint_app_name"] = tablePinpointAppName
	primaryKeyMap["pinpoint_app_name"] = tablePinpointAppName.PrimaryKey()
	//187 plan_confirm
	tablePlanConfirm := PlanConfirm{}
	modelMap["plan_confirm"] = tablePlanConfirm
	primaryKeyMap["plan_confirm"] = tablePlanConfirm.PrimaryKey()
	//188 plan_jira_projects
	tablePlanJiraProjects := PlanJiraProjects{}
	modelMap["plan_jira_projects"] = tablePlanJiraProjects
	primaryKeyMap["plan_jira_projects"] = tablePlanJiraProjects.PrimaryKey()
	//189 plan_log
	tablePlanLog := PlanLog{}
	modelMap["plan_log"] = tablePlanLog
	primaryKeyMap["plan_log"] = tablePlanLog.PrimaryKey()
	//190 plan_notice
	tablePlanNotice := PlanNotice{}
	modelMap["plan_notice"] = tablePlanNotice
	primaryKeyMap["plan_notice"] = tablePlanNotice.PrimaryKey()
	//191 plan_role_users
	tablePlanRoleUsers := PlanRoleUsers{}
	modelMap["plan_role_users"] = tablePlanRoleUsers
	primaryKeyMap["plan_role_users"] = tablePlanRoleUsers.PrimaryKey()
	//192 product_line
	tableProductLine := ProductLine{}
	modelMap["product_line"] = tableProductLine
	primaryKeyMap["product_line"] = tableProductLine.PrimaryKey()
	//193 project_product_relation
	tableProjectProductRelation := ProjectProductRelation{}
	modelMap["project_product_relation"] = tableProjectProductRelation
	primaryKeyMap["project_product_relation"] = tableProjectProductRelation.PrimaryKey()
	//194 pubfile_agg_issue_list
	tablePubfileAggIssueList := PubfileAggIssueList{}
	modelMap["pubfile_agg_issue_list"] = tablePubfileAggIssueList
	primaryKeyMap["pubfile_agg_issue_list"] = tablePubfileAggIssueList.PrimaryKey()
	//195 pubfile_aggregation_control
	tablePubfileAggregationControl := PubfileAggregationControl{}
	modelMap["pubfile_aggregation_control"] = tablePubfileAggregationControl
	primaryKeyMap["pubfile_aggregation_control"] = tablePubfileAggregationControl.PrimaryKey()
	//196 pubfile_aggregation_rel
	tablePubfileAggregationRel := PubfileAggregationRel{}
	modelMap["pubfile_aggregation_rel"] = tablePubfileAggregationRel
	primaryKeyMap["pubfile_aggregation_rel"] = tablePubfileAggregationRel.PrimaryKey()
	//197 pubfile_aggregation_rel_req
	tablePubfileAggregationRelReq := PubfileAggregationRelReq{}
	modelMap["pubfile_aggregation_rel_req"] = tablePubfileAggregationRelReq
	primaryKeyMap["pubfile_aggregation_rel_req"] = tablePubfileAggregationRelReq.PrimaryKey()
	//198 pubfile_version_control
	tablePubfileVersionControl := PubfileVersionControl{}
	modelMap["pubfile_version_control"] = tablePubfileVersionControl
	primaryKeyMap["pubfile_version_control"] = tablePubfileVersionControl.PrimaryKey()
	//199 publish_issue_list
	tablePublishIssueList := PublishIssueList{}
	modelMap["publish_issue_list"] = tablePublishIssueList
	primaryKeyMap["publish_issue_list"] = tablePublishIssueList.PrimaryKey()
	//200 publish_notice
	tablePublishNotice := PublishNotice{}
	modelMap["publish_notice"] = tablePublishNotice
	primaryKeyMap["publish_notice"] = tablePublishNotice.PrimaryKey()
	//201 publish_notice_type
	tablePublishNoticeType := PublishNoticeType{}
	modelMap["publish_notice_type"] = tablePublishNoticeType
	primaryKeyMap["publish_notice_type"] = tablePublishNoticeType.PrimaryKey()
	//202 publish_pannel
	tablePublishPannel := PublishPannel{}
	modelMap["publish_pannel"] = tablePublishPannel
	primaryKeyMap["publish_pannel"] = tablePublishPannel.PrimaryKey()
	//203 publish_redis_cache
	tablePublishRedisCache := PublishRedisCache{}
	modelMap["publish_redis_cache"] = tablePublishRedisCache
	primaryKeyMap["publish_redis_cache"] = tablePublishRedisCache.PrimaryKey()
	//204 publish_req
	tablePublishReq := PublishReq{}
	modelMap["publish_req"] = tablePublishReq
	primaryKeyMap["publish_req"] = tablePublishReq.PrimaryKey()
	//205 publish_req_before
	tablePublishReqBefore := PublishReqBefore{}
	modelMap["publish_req_before"] = tablePublishReqBefore
	primaryKeyMap["publish_req_before"] = tablePublishReqBefore.PrimaryKey()
	//206 publish_req_copy
	tablePublishReqCopy := PublishReqCopy{}
	modelMap["publish_req_copy"] = tablePublishReqCopy
	primaryKeyMap["publish_req_copy"] = tablePublishReqCopy.PrimaryKey()
	//207 publish_req_dbscript
	tablePublishReqDbscript := PublishReqDbscript{}
	modelMap["publish_req_dbscript"] = tablePublishReqDbscript
	primaryKeyMap["publish_req_dbscript"] = tablePublishReqDbscript.PrimaryKey()
	//208 publish_req_deploy_event
	tablePublishReqDeployEvent := PublishReqDeployEvent{}
	modelMap["publish_req_deploy_event"] = tablePublishReqDeployEvent
	primaryKeyMap["publish_req_deploy_event"] = tablePublishReqDeployEvent.PrimaryKey()
	//209 publish_req_deploy_product_env
	tablePublishReqDeployProductEnv := PublishReqDeployProductEnv{}
	modelMap["publish_req_deploy_product_env"] = tablePublishReqDeployProductEnv
	primaryKeyMap["publish_req_deploy_product_env"] = tablePublishReqDeployProductEnv.PrimaryKey()
	//210 publish_req_details
	tablePublishReqDetails := PublishReqDetails{}
	modelMap["publish_req_details"] = tablePublishReqDetails
	primaryKeyMap["publish_req_details"] = tablePublishReqDetails.PrimaryKey()
	//211 publish_req_env
	tablePublishReqEnv := PublishReqEnv{}
	modelMap["publish_req_env"] = tablePublishReqEnv
	primaryKeyMap["publish_req_env"] = tablePublishReqEnv.PrimaryKey()
	//212 publish_req_menu_script
	tablePublishReqMenuScript := PublishReqMenuScript{}
	modelMap["publish_req_menu_script"] = tablePublishReqMenuScript
	primaryKeyMap["publish_req_menu_script"] = tablePublishReqMenuScript.PrimaryKey()
	//213 publish_req_plan
	tablePublishReqPlan := PublishReqPlan{}
	modelMap["publish_req_plan"] = tablePublishReqPlan
	primaryKeyMap["publish_req_plan"] = tablePublishReqPlan.PrimaryKey()
	//214 publish_req_rollback_event
	tablePublishReqRollbackEvent := PublishReqRollbackEvent{}
	modelMap["publish_req_rollback_event"] = tablePublishReqRollbackEvent
	primaryKeyMap["publish_req_rollback_event"] = tablePublishReqRollbackEvent.PrimaryKey()
	//215 publish_req_statu
	tablePublishReqStatu := PublishReqStatu{}
	modelMap["publish_req_statu"] = tablePublishReqStatu
	primaryKeyMap["publish_req_statu"] = tablePublishReqStatu.PrimaryKey()
	//216 publish_req_status_change_reason
	tablePublishReqStatusChangeReason := PublishReqStatusChangeReason{}
	modelMap["publish_req_status_change_reason"] = tablePublishReqStatusChangeReason
	primaryKeyMap["publish_req_status_change_reason"] = tablePublishReqStatusChangeReason.PrimaryKey()
	//217 publish_req_version
	tablePublishReqVersion := PublishReqVersion{}
	modelMap["publish_req_version"] = tablePublishReqVersion
	primaryKeyMap["publish_req_version"] = tablePublishReqVersion.PrimaryKey()
	//218 publish_role_users
	tablePublishRoleUsers := PublishRoleUsers{}
	modelMap["publish_role_users"] = tablePublishRoleUsers
	primaryKeyMap["publish_role_users"] = tablePublishRoleUsers.PrimaryKey()
	//219 publish_wiki_list
	tablePublishWikiList := PublishWikiList{}
	modelMap["publish_wiki_list"] = tablePublishWikiList
	primaryKeyMap["publish_wiki_list"] = tablePublishWikiList.PrimaryKey()
	//220 record_care
	tableRecordCare := RecordCare{}
	modelMap["record_care"] = tableRecordCare
	primaryKeyMap["record_care"] = tableRecordCare.PrimaryKey()
	//221 record_care_inner
	tableRecordCareInner := RecordCareInner{}
	modelMap["record_care_inner"] = tableRecordCareInner
	primaryKeyMap["record_care_inner"] = tableRecordCareInner.PrimaryKey()
	//222 record_comment
	tableRecordComment := RecordComment{}
	modelMap["record_comment"] = tableRecordComment
	primaryKeyMap["record_comment"] = tableRecordComment.PrimaryKey()
	//223 record_issues
	tableRecordIssues := RecordIssues{}
	modelMap["record_issues"] = tableRecordIssues
	primaryKeyMap["record_issues"] = tableRecordIssues.PrimaryKey()
	//224 recovery_impl
	tableRecoveryImpl := RecoveryImpl{}
	modelMap["recovery_impl"] = tableRecoveryImpl
	primaryKeyMap["recovery_impl"] = tableRecoveryImpl.PrimaryKey()
	//225 recovery_impl_package
	tableRecoveryImplPackage := RecoveryImplPackage{}
	modelMap["recovery_impl_package"] = tableRecoveryImplPackage
	primaryKeyMap["recovery_impl_package"] = tableRecoveryImplPackage.PrimaryKey()
	//226 recovery_impl_rel
	tableRecoveryImplRel := RecoveryImplRel{}
	modelMap["recovery_impl_rel"] = tableRecoveryImplRel
	primaryKeyMap["recovery_impl_rel"] = tableRecoveryImplRel.PrimaryKey()
	//227 recovery_impl_response
	tableRecoveryImplResponse := RecoveryImplResponse{}
	modelMap["recovery_impl_response"] = tableRecoveryImplResponse
	primaryKeyMap["recovery_impl_response"] = tableRecoveryImplResponse.PrimaryKey()
	//228 recovery_strategy
	tableRecoveryStrategy := RecoveryStrategy{}
	modelMap["recovery_strategy"] = tableRecoveryStrategy
	primaryKeyMap["recovery_strategy"] = tableRecoveryStrategy.PrimaryKey()
	//229 req_deploy_tag
	tableReqDeployTag := ReqDeployTag{}
	modelMap["req_deploy_tag"] = tableReqDeployTag
	primaryKeyMap["req_deploy_tag"] = tableReqDeployTag.PrimaryKey()
	//230 req_info_modify_detail
	tableReqInfoModifyDetail := ReqInfoModifyDetail{}
	modelMap["req_info_modify_detail"] = tableReqInfoModifyDetail
	primaryKeyMap["req_info_modify_detail"] = tableReqInfoModifyDetail.PrimaryKey()
	//231 req_info_modify_trace
	tableReqInfoModifyTrace := ReqInfoModifyTrace{}
	modelMap["req_info_modify_trace"] = tableReqInfoModifyTrace
	primaryKeyMap["req_info_modify_trace"] = tableReqInfoModifyTrace.PrimaryKey()
	//232 req_issue_record
	tableReqIssueRecord := ReqIssueRecord{}
	modelMap["req_issue_record"] = tableReqIssueRecord
	primaryKeyMap["req_issue_record"] = tableReqIssueRecord.PrimaryKey()
	//233 req_issue_record_img
	tableReqIssueRecordImg := ReqIssueRecordImg{}
	modelMap["req_issue_record_img"] = tableReqIssueRecordImg
	primaryKeyMap["req_issue_record_img"] = tableReqIssueRecordImg.PrimaryKey()
	//234 req_notice_about
	tableReqNoticeAbout := ReqNoticeAbout{}
	modelMap["req_notice_about"] = tableReqNoticeAbout
	primaryKeyMap["req_notice_about"] = tableReqNoticeAbout.PrimaryKey()
	//235 role_ldap_group
	tableRoleLdapGroup := RoleLdapGroup{}
	modelMap["role_ldap_group"] = tableRoleLdapGroup
	primaryKeyMap["role_ldap_group"] = tableRoleLdapGroup.PrimaryKey()
	//236 script_exe_blacklist
	tableScriptExeBlacklist := ScriptExeBlacklist{}
	modelMap["script_exe_blacklist"] = tableScriptExeBlacklist
	primaryKeyMap["script_exe_blacklist"] = tableScriptExeBlacklist.PrimaryKey()
	//237 sentry_issues
	tableSentryIssues := SentryIssues{}
	modelMap["sentry_issues"] = tableSentryIssues
	primaryKeyMap["sentry_issues"] = tableSentryIssues.PrimaryKey()
	//238 sentry_record_issues
	tableSentryRecordIssues := SentryRecordIssues{}
	modelMap["sentry_record_issues"] = tableSentryRecordIssues
	primaryKeyMap["sentry_record_issues"] = tableSentryRecordIssues.PrimaryKey()
	//239 server_app_log
	tableServerAppLog := ServerAppLog{}
	modelMap["server_app_log"] = tableServerAppLog
	primaryKeyMap["server_app_log"] = tableServerAppLog.PrimaryKey()
	//240 server_dbsync_detail
	tableServerDbsyncDetail := ServerDbsyncDetail{}
	modelMap["server_dbsync_detail"] = tableServerDbsyncDetail
	primaryKeyMap["server_dbsync_detail"] = tableServerDbsyncDetail.PrimaryKey()
	//241 server_env
	tableServerEnv := ServerEnv{}
	modelMap["server_env"] = tableServerEnv
	primaryKeyMap["server_env"] = tableServerEnv.PrimaryKey()
	//242 server_group
	tableServerGroup := ServerGroup{}
	modelMap["server_group"] = tableServerGroup
	primaryKeyMap["server_group"] = tableServerGroup.PrimaryKey()
	//243 server_hardware
	tableServerHardware := ServerHardware{}
	modelMap["server_hardware"] = tableServerHardware
	primaryKeyMap["server_hardware"] = tableServerHardware.PrimaryKey()
	//244 server_hardware_copy
	tableServerHardwareCopy := ServerHardwareCopy{}
	modelMap["server_hardware_copy"] = tableServerHardwareCopy
	primaryKeyMap["server_hardware_copy"] = tableServerHardwareCopy.PrimaryKey()
	//245 server_hardware_copy2
	tableServerHardwareCopy2 := ServerHardwareCopy2{}
	modelMap["server_hardware_copy2"] = tableServerHardwareCopy2
	primaryKeyMap["server_hardware_copy2"] = tableServerHardwareCopy2.PrimaryKey()
	//246 server_install_soft
	tableServerInstallSoft := ServerInstallSoft{}
	modelMap["server_install_soft"] = tableServerInstallSoft
	primaryKeyMap["server_install_soft"] = tableServerInstallSoft.PrimaryKey()
	//247 server_install_soft_template
	tableServerInstallSoftTemplate := ServerInstallSoftTemplate{}
	modelMap["server_install_soft_template"] = tableServerInstallSoftTemplate
	primaryKeyMap["server_install_soft_template"] = tableServerInstallSoftTemplate.PrimaryKey()
	//248 server_instance
	tableServerInstance := ServerInstance{}
	modelMap["server_instance"] = tableServerInstance
	primaryKeyMap["server_instance"] = tableServerInstance.PrimaryKey()
	//249 server_instance_copy
	tableServerInstanceCopy := ServerInstanceCopy{}
	modelMap["server_instance_copy"] = tableServerInstanceCopy
	primaryKeyMap["server_instance_copy"] = tableServerInstanceCopy.PrimaryKey()
	//250 server_labels
	tableServerLabels := ServerLabels{}
	modelMap["server_labels"] = tableServerLabels
	primaryKeyMap["server_labels"] = tableServerLabels.PrimaryKey()
	//251 server_modify_recoder
	tableServerModifyRecoder := ServerModifyRecoder{}
	modelMap["server_modify_recoder"] = tableServerModifyRecoder
	primaryKeyMap["server_modify_recoder"] = tableServerModifyRecoder.PrimaryKey()
	//252 server_ops_log
	tableServerOpsLog := ServerOpsLog{}
	modelMap["server_ops_log"] = tableServerOpsLog
	primaryKeyMap["server_ops_log"] = tableServerOpsLog.PrimaryKey()
	//253 server_ops_log_detail
	tableServerOpsLogDetail := ServerOpsLogDetail{}
	modelMap["server_ops_log_detail"] = tableServerOpsLogDetail
	primaryKeyMap["server_ops_log_detail"] = tableServerOpsLogDetail.PrimaryKey()
	//254 server_os_users
	tableServerOsUsers := ServerOsUsers{}
	modelMap["server_os_users"] = tableServerOsUsers
	primaryKeyMap["server_os_users"] = tableServerOsUsers.PrimaryKey()
	//255 server_service_type
	tableServerServiceType := ServerServiceType{}
	modelMap["server_service_type"] = tableServerServiceType
	primaryKeyMap["server_service_type"] = tableServerServiceType.PrimaryKey()
	//256 server_soft_conf
	tableServerSoftConf := ServerSoftConf{}
	modelMap["server_soft_conf"] = tableServerSoftConf
	primaryKeyMap["server_soft_conf"] = tableServerSoftConf.PrimaryKey()
	//257 server_soft_conf_def
	tableServerSoftConfDef := ServerSoftConfDef{}
	modelMap["server_soft_conf_def"] = tableServerSoftConfDef
	primaryKeyMap["server_soft_conf_def"] = tableServerSoftConfDef.PrimaryKey()
	//258 server_soft_install_event
	tableServerSoftInstallEvent := ServerSoftInstallEvent{}
	modelMap["server_soft_install_event"] = tableServerSoftInstallEvent
	primaryKeyMap["server_soft_install_event"] = tableServerSoftInstallEvent.PrimaryKey()
	//259 server_soft_logs
	tableServerSoftLogs := ServerSoftLogs{}
	modelMap["server_soft_logs"] = tableServerSoftLogs
	primaryKeyMap["server_soft_logs"] = tableServerSoftLogs.PrimaryKey()
	//260 server_soft_user
	tableServerSoftUser := ServerSoftUser{}
	modelMap["server_soft_user"] = tableServerSoftUser
	primaryKeyMap["server_soft_user"] = tableServerSoftUser.PrimaryKey()
	//261 server_soft_user_group_rolerel
	tableServerSoftUserGroupRolerel := ServerSoftUserGroupRolerel{}
	modelMap["server_soft_user_group_rolerel"] = tableServerSoftUserGroupRolerel
	primaryKeyMap["server_soft_user_group_rolerel"] = tableServerSoftUserGroupRolerel.PrimaryKey()
	//262 server_ssh_event
	tableServerSSHEvent := ServerSSHEvent{}
	modelMap["server_ssh_event"] = tableServerSSHEvent
	primaryKeyMap["server_ssh_event"] = tableServerSSHEvent.PrimaryKey()
	//263 server_vps_event
	tableServerVpsEvent := ServerVpsEvent{}
	modelMap["server_vps_event"] = tableServerVpsEvent
	primaryKeyMap["server_vps_event"] = tableServerVpsEvent.PrimaryKey()
	//264 server_vsphere
	tableServerVsphere := ServerVsphere{}
	modelMap["server_vsphere"] = tableServerVsphere
	primaryKeyMap["server_vsphere"] = tableServerVsphere.PrimaryKey()
	//265 single_file_depbak
	tableSingleFileDepbak := SingleFileDepbak{}
	modelMap["single_file_depbak"] = tableSingleFileDepbak
	primaryKeyMap["single_file_depbak"] = tableSingleFileDepbak.PrimaryKey()
	//266 soft_cmd_list
	tableSoftCmdList := SoftCmdList{}
	modelMap["soft_cmd_list"] = tableSoftCmdList
	primaryKeyMap["soft_cmd_list"] = tableSoftCmdList.PrimaryKey()
	//267 soft_cmd_type
	tableSoftCmdType := SoftCmdType{}
	modelMap["soft_cmd_type"] = tableSoftCmdType
	primaryKeyMap["soft_cmd_type"] = tableSoftCmdType.PrimaryKey()
	//268 soft_logs
	tableSoftLogs := SoftLogs{}
	modelMap["soft_logs"] = tableSoftLogs
	primaryKeyMap["soft_logs"] = tableSoftLogs.PrimaryKey()
	//269 soft_model_new
	tableSoftModelNew := SoftModelNew{}
	modelMap["soft_model_new"] = tableSoftModelNew
	primaryKeyMap["soft_model_new"] = tableSoftModelNew.PrimaryKey()
	//270 soft_slave_ip
	tableSoftSlaveIP := SoftSlaveIP{}
	modelMap["soft_slave_ip"] = tableSoftSlaveIP
	primaryKeyMap["soft_slave_ip"] = tableSoftSlaveIP.PrimaryKey()
	//271 sql_admin_api_token
	tableSqlAdminAPIToken := SqlAdminAPIToken{}
	modelMap["sql_admin_api_token"] = tableSqlAdminAPIToken
	primaryKeyMap["sql_admin_api_token"] = tableSqlAdminAPIToken.PrimaryKey()
	//272 sql_admin_auth
	tableSqlAdminAuth := SqlAdminAuth{}
	modelMap["sql_admin_auth"] = tableSqlAdminAuth
	primaryKeyMap["sql_admin_auth"] = tableSqlAdminAuth.PrimaryKey()
	//273 sql_admin_favorite
	tableSqlAdminFavorite := SqlAdminFavorite{}
	modelMap["sql_admin_favorite"] = tableSqlAdminFavorite
	primaryKeyMap["sql_admin_favorite"] = tableSqlAdminFavorite.PrimaryKey()
	//274 sql_admin_log
	tableSqlAdminLog := SqlAdminLog{}
	modelMap["sql_admin_log"] = tableSqlAdminLog
	primaryKeyMap["sql_admin_log"] = tableSqlAdminLog.PrimaryKey()
	//275 sql_admin_not_commit
	tableSqlAdminNotCommit := SqlAdminNotCommit{}
	modelMap["sql_admin_not_commit"] = tableSqlAdminNotCommit
	primaryKeyMap["sql_admin_not_commit"] = tableSqlAdminNotCommit.PrimaryKey()
	//276 sql_admin_role
	tableSqlAdminRole := SqlAdminRole{}
	modelMap["sql_admin_role"] = tableSqlAdminRole
	primaryKeyMap["sql_admin_role"] = tableSqlAdminRole.PrimaryKey()
	//277 sql_admin_snippets
	tableSqlAdminSnippets := SqlAdminSnippets{}
	modelMap["sql_admin_snippets"] = tableSqlAdminSnippets
	primaryKeyMap["sql_admin_snippets"] = tableSqlAdminSnippets.PrimaryKey()
	//278 sql_admin_websqlrule
	tableSqlAdminWebsqlrule := SqlAdminWebsqlrule{}
	modelMap["sql_admin_websqlrule"] = tableSqlAdminWebsqlrule
	primaryKeyMap["sql_admin_websqlrule"] = tableSqlAdminWebsqlrule.PrimaryKey()
	//279 sql_script
	tableSqlScript := SqlScript{}
	modelMap["sql_script"] = tableSqlScript
	primaryKeyMap["sql_script"] = tableSqlScript.PrimaryKey()
	//280 sql_script_label_relation
	tableSqlScriptLabelRelation := SqlScriptLabelRelation{}
	modelMap["sql_script_label_relation"] = tableSqlScriptLabelRelation
	primaryKeyMap["sql_script_label_relation"] = tableSqlScriptLabelRelation.PrimaryKey()
	//281 sql_script_link
	tableSqlScriptLink := SqlScriptLink{}
	modelMap["sql_script_link"] = tableSqlScriptLink
	primaryKeyMap["sql_script_link"] = tableSqlScriptLink.PrimaryKey()
	//282 stability_assurance_jdump
	tableStabilityAssuranceJdump := StabilityAssuranceJdump{}
	modelMap["stability_assurance_jdump"] = tableStabilityAssuranceJdump
	primaryKeyMap["stability_assurance_jdump"] = tableStabilityAssuranceJdump.PrimaryKey()
	//283 sys_board
	tableSysBoard := SysBoard{}
	modelMap["sys_board"] = tableSysBoard
	primaryKeyMap["sys_board"] = tableSysBoard.PrimaryKey()
	//284 sys_cmd
	tableSysCmd := SysCmd{}
	modelMap["sys_cmd"] = tableSysCmd
	primaryKeyMap["sys_cmd"] = tableSysCmd.PrimaryKey()
	//285 sys_code
	tableSysCode := SysCode{}
	modelMap["sys_code"] = tableSysCode
	primaryKeyMap["sys_code"] = tableSysCode.PrimaryKey()
	//286 sys_common_config
	tableSysCommonConfig := SysCommonConfig{}
	modelMap["sys_common_config"] = tableSysCommonConfig
	primaryKeyMap["sys_common_config"] = tableSysCommonConfig.PrimaryKey()
	//287 sys_dbuser
	tableSysDbuser := SysDbuser{}
	modelMap["sys_dbuser"] = tableSysDbuser
	primaryKeyMap["sys_dbuser"] = tableSysDbuser.PrimaryKey()
	//288 sys_env
	tableSysEnv := SysEnv{}
	modelMap["sys_env"] = tableSysEnv
	primaryKeyMap["sys_env"] = tableSysEnv.PrimaryKey()
	//289 sys_env_param
	tableSysEnvParam := SysEnvParam{}
	modelMap["sys_env_param"] = tableSysEnvParam
	primaryKeyMap["sys_env_param"] = tableSysEnvParam.PrimaryKey()
	//290 sys_error_code
	tableSysErrorCode := SysErrorCode{}
	modelMap["sys_error_code"] = tableSysErrorCode
	primaryKeyMap["sys_error_code"] = tableSysErrorCode.PrimaryKey()
	//291 sys_faq
	tableSysFaq := SysFaq{}
	modelMap["sys_faq"] = tableSysFaq
	primaryKeyMap["sys_faq"] = tableSysFaq.PrimaryKey()
	//292 sys_form
	tableSysForm := SysForm{}
	modelMap["sys_form"] = tableSysForm
	primaryKeyMap["sys_form"] = tableSysForm.PrimaryKey()
	//293 sys_form_fields
	tableSysFormFields := SysFormFields{}
	modelMap["sys_form_fields"] = tableSysFormFields
	primaryKeyMap["sys_form_fields"] = tableSysFormFields.PrimaryKey()
	//294 sys_func
	tableSysFunc := SysFunc{}
	modelMap["sys_func"] = tableSysFunc
	primaryKeyMap["sys_func"] = tableSysFunc.PrimaryKey()
	//295 sys_hospital
	tableSysHospital := SysHospital{}
	modelMap["sys_hospital"] = tableSysHospital
	primaryKeyMap["sys_hospital"] = tableSysHospital.PrimaryKey()
	//296 sys_host
	tableSysHost := SysHost{}
	modelMap["sys_host"] = tableSysHost
	primaryKeyMap["sys_host"] = tableSysHost.PrimaryKey()
	//297 sys_info
	tableSysInfo := SysInfo{}
	modelMap["sys_info"] = tableSysInfo
	primaryKeyMap["sys_info"] = tableSysInfo.PrimaryKey()
	//298 sys_labels
	tableSysLabels := SysLabels{}
	modelMap["sys_labels"] = tableSysLabels
	primaryKeyMap["sys_labels"] = tableSysLabels.PrimaryKey()
	//299 sys_like
	tableSysLike := SysLike{}
	modelMap["sys_like"] = tableSysLike
	primaryKeyMap["sys_like"] = tableSysLike.PrimaryKey()
	//300 sys_menu
	tableSysMenu := SysMenu{}
	modelMap["sys_menu"] = tableSysMenu
	primaryKeyMap["sys_menu"] = tableSysMenu.PrimaryKey()
	//301 sys_menu_v2
	tableSysMenuV2 := SysMenuV2{}
	modelMap["sys_menu_v2"] = tableSysMenuV2
	primaryKeyMap["sys_menu_v2"] = tableSysMenuV2.PrimaryKey()
	//302 sys_param
	tableSysParam := SysParam{}
	modelMap["sys_param"] = tableSysParam
	primaryKeyMap["sys_param"] = tableSysParam.PrimaryKey()
	//303 sys_param_copy
	tableSysParamCopy := SysParamCopy{}
	modelMap["sys_param_copy"] = tableSysParamCopy
	primaryKeyMap["sys_param_copy"] = tableSysParamCopy.PrimaryKey()
	//304 sys_param_copy1
	tableSysParamCopy1 := SysParamCopy1{}
	modelMap["sys_param_copy1"] = tableSysParamCopy1
	primaryKeyMap["sys_param_copy1"] = tableSysParamCopy1.PrimaryKey()
	//305 sys_portal_type
	tableSysPortalType := SysPortalType{}
	modelMap["sys_portal_type"] = tableSysPortalType
	primaryKeyMap["sys_portal_type"] = tableSysPortalType.PrimaryKey()
	//306 sys_portals
	tableSysPortals := SysPortals{}
	modelMap["sys_portals"] = tableSysPortals
	primaryKeyMap["sys_portals"] = tableSysPortals.PrimaryKey()
	//307 sys_project_param
	tableSysProjectParam := SysProjectParam{}
	modelMap["sys_project_param"] = tableSysProjectParam
	primaryKeyMap["sys_project_param"] = tableSysProjectParam.PrimaryKey()
	//308 sys_project_team
	tableSysProjectTeam := SysProjectTeam{}
	modelMap["sys_project_team"] = tableSysProjectTeam
	primaryKeyMap["sys_project_team"] = tableSysProjectTeam.PrimaryKey()
	//309 sys_project_team_leader
	tableSysProjectTeamLeader := SysProjectTeamLeader{}
	modelMap["sys_project_team_leader"] = tableSysProjectTeamLeader
	primaryKeyMap["sys_project_team_leader"] = tableSysProjectTeamLeader.PrimaryKey()
	//310 sys_projects
	tableSysProjects := SysProjects{}
	modelMap["sys_projects"] = tableSysProjects
	primaryKeyMap["sys_projects"] = tableSysProjects.PrimaryKey()
	//311 sys_server_type
	tableSysServerType := SysServerType{}
	modelMap["sys_server_type"] = tableSysServerType
	primaryKeyMap["sys_server_type"] = tableSysServerType.PrimaryKey()
	//312 sys_template
	tableSysTemplate := SysTemplate{}
	modelMap["sys_template"] = tableSysTemplate
	primaryKeyMap["sys_template"] = tableSysTemplate.PrimaryKey()
	//313 sys_token
	tableSysToken := SysToken{}
	modelMap["sys_token"] = tableSysToken
	primaryKeyMap["sys_token"] = tableSysToken.PrimaryKey()
	//314 sys_token_business
	tableSysTokenBusiness := SysTokenBusiness{}
	modelMap["sys_token_business"] = tableSysTokenBusiness
	primaryKeyMap["sys_token_business"] = tableSysTokenBusiness.PrimaryKey()
	//315 sys_token_user
	tableSysTokenUser := SysTokenUser{}
	modelMap["sys_token_user"] = tableSysTokenUser
	primaryKeyMap["sys_token_user"] = tableSysTokenUser.PrimaryKey()
	//316 sys_token_user_group
	tableSysTokenUserGroup := SysTokenUserGroup{}
	modelMap["sys_token_user_group"] = tableSysTokenUserGroup
	primaryKeyMap["sys_token_user_group"] = tableSysTokenUserGroup.PrimaryKey()
	//317 sys_uum_system
	tableSysUumSystem := SysUumSystem{}
	modelMap["sys_uum_system"] = tableSysUumSystem
	primaryKeyMap["sys_uum_system"] = tableSysUumSystem.PrimaryKey()
	//318 sys_war
	tableSysWar := SysWar{}
	modelMap["sys_war"] = tableSysWar
	primaryKeyMap["sys_war"] = tableSysWar.PrimaryKey()
	//319 sys_war_731
	tableSysWar731 := SysWar731{}
	modelMap["sys_war_731"] = tableSysWar731
	primaryKeyMap["sys_war_731"] = tableSysWar731.PrimaryKey()
	//320 sys_war_name
	tableSysWarName := SysWarName{}
	modelMap["sys_war_name"] = tableSysWarName
	primaryKeyMap["sys_war_name"] = tableSysWarName.PrimaryKey()
	//321 sys_website
	tableSysWebsite := SysWebsite{}
	modelMap["sys_website"] = tableSysWebsite
	primaryKeyMap["sys_website"] = tableSysWebsite.PrimaryKey()
	//322 t1
	tableT1 := T1{}
	modelMap["t1"] = tableT1
	primaryKeyMap["t1"] = tableT1.PrimaryKey()
	//323 tag_desc
	tableTagDesc := TagDesc{}
	modelMap["tag_desc"] = tableTagDesc
	primaryKeyMap["tag_desc"] = tableTagDesc.PrimaryKey()
	//324 template_event_log
	tableTemplateEventLog := TemplateEventLog{}
	modelMap["template_event_log"] = tableTemplateEventLog
	primaryKeyMap["template_event_log"] = tableTemplateEventLog.PrimaryKey()
	//325 tmp
	tableTmp := Tmp{}
	modelMap["tmp"] = tableTmp
	primaryKeyMap["tmp"] = tableTmp.PrimaryKey()
	//326 topography_map
	tableTopographyMap := TopographyMap{}
	modelMap["topography_map"] = tableTopographyMap
	primaryKeyMap["topography_map"] = tableTopographyMap.PrimaryKey()
	//327 ue_planning
	tableUePlanning := UePlanning{}
	modelMap["ue_planning"] = tableUePlanning
	primaryKeyMap["ue_planning"] = tableUePlanning.PrimaryKey()
	//328 ueplain_issues
	tableUeplainIssues := UeplainIssues{}
	modelMap["ueplain_issues"] = tableUeplainIssues
	primaryKeyMap["ueplain_issues"] = tableUeplainIssues.PrimaryKey()
	//329 upgrade_plan_dbscript
	tableUpgradePlanDbscript := UpgradePlanDbscript{}
	modelMap["upgrade_plan_dbscript"] = tableUpgradePlanDbscript
	primaryKeyMap["upgrade_plan_dbscript"] = tableUpgradePlanDbscript.PrimaryKey()
	//330 upgrade_plan_deatils
	tableUpgradePlanDeatils := UpgradePlanDeatils{}
	modelMap["upgrade_plan_deatils"] = tableUpgradePlanDeatils
	primaryKeyMap["upgrade_plan_deatils"] = tableUpgradePlanDeatils.PrimaryKey()
	//331 upgrade_plan_list
	tableUpgradePlanList := UpgradePlanList{}
	modelMap["upgrade_plan_list"] = tableUpgradePlanList
	primaryKeyMap["upgrade_plan_list"] = tableUpgradePlanList.PrimaryKey()
	//332 user
	tableUser := User{}
	modelMap["user"] = tableUser
	primaryKeyMap["user"] = tableUser.PrimaryKey()
	//333 user_account
	tableUserAccount := UserAccount{}
	modelMap["user_account"] = tableUserAccount
	primaryKeyMap["user_account"] = tableUserAccount.PrimaryKey()
	//334 user_account_bak
	tableUserAccountBak := UserAccountBak{}
	modelMap["user_account_bak"] = tableUserAccountBak
	primaryKeyMap["user_account_bak"] = tableUserAccountBak.PrimaryKey()
	//335 user_account_copy
	tableUserAccountCopy := UserAccountCopy{}
	modelMap["user_account_copy"] = tableUserAccountCopy
	primaryKeyMap["user_account_copy"] = tableUserAccountCopy.PrimaryKey()
	//336 user_api_key
	tableUserAPIKey := UserAPIKey{}
	modelMap["user_api_key"] = tableUserAPIKey
	primaryKeyMap["user_api_key"] = tableUserAPIKey.PrimaryKey()
	//337 user_favorites
	tableUserFavorites := UserFavorites{}
	modelMap["user_favorites"] = tableUserFavorites
	primaryKeyMap["user_favorites"] = tableUserFavorites.PrimaryKey()
	//338 user_group
	tableUserGroup := UserGroup{}
	modelMap["user_group"] = tableUserGroup
	primaryKeyMap["user_group"] = tableUserGroup.PrimaryKey()
	//339 user_group_rel
	tableUserGroupRel := UserGroupRel{}
	modelMap["user_group_rel"] = tableUserGroupRel
	primaryKeyMap["user_group_rel"] = tableUserGroupRel.PrimaryKey()
	//340 user_group_type
	tableUserGroupType := UserGroupType{}
	modelMap["user_group_type"] = tableUserGroupType
	primaryKeyMap["user_group_type"] = tableUserGroupType.PrimaryKey()
	//341 user_role
	tableUserRole := UserRole{}
	modelMap["user_role"] = tableUserRole
	primaryKeyMap["user_role"] = tableUserRole.PrimaryKey()
	//342 user_role_env
	tableUserRoleEnv := UserRoleEnv{}
	modelMap["user_role_env"] = tableUserRoleEnv
	primaryKeyMap["user_role_env"] = tableUserRoleEnv.PrimaryKey()
	//343 user_role_func
	tableUserRoleFunc := UserRoleFunc{}
	modelMap["user_role_func"] = tableUserRoleFunc
	primaryKeyMap["user_role_func"] = tableUserRoleFunc.PrimaryKey()
	//344 user_role_hospital
	tableUserRoleHospital := UserRoleHospital{}
	modelMap["user_role_hospital"] = tableUserRoleHospital
	primaryKeyMap["user_role_hospital"] = tableUserRoleHospital.PrimaryKey()
	//345 user_role_id_list
	tableUserRoleIDList := UserRoleIDList{}
	modelMap["user_role_id_list"] = tableUserRoleIDList
	primaryKeyMap["user_role_id_list"] = tableUserRoleIDList.PrimaryKey()
	//346 user_role_menu
	tableUserRoleMenu := UserRoleMenu{}
	modelMap["user_role_menu"] = tableUserRoleMenu
	primaryKeyMap["user_role_menu"] = tableUserRoleMenu.PrimaryKey()
	//347 user_tag
	tableUserTag := UserTag{}
	modelMap["user_tag"] = tableUserTag
	primaryKeyMap["user_tag"] = tableUserTag.PrimaryKey()
	//348 users
	tableUsers := Users{}
	modelMap["users"] = tableUsers
	primaryKeyMap["users"] = tableUsers.PrimaryKey()
	//349 users_groups
	tableUsersGroups := UsersGroups{}
	modelMap["users_groups"] = tableUsersGroups
	primaryKeyMap["users_groups"] = tableUsersGroups.PrimaryKey()
	//350 users_permissions
	tableUsersPermissions := UsersPermissions{}
	modelMap["users_permissions"] = tableUsersPermissions
	primaryKeyMap["users_permissions"] = tableUsersPermissions.PrimaryKey()
	//351 version_plan
	tableVersionPlan := VersionPlan{}
	modelMap["version_plan"] = tableVersionPlan
	primaryKeyMap["version_plan"] = tableVersionPlan.PrimaryKey()
	//352 view_server_env_list
	tableViewServerEnvList := ViewServerEnvList{}
	modelMap["view_server_env_list"] = tableViewServerEnvList
	primaryKeyMap["view_server_env_list"] = tableViewServerEnvList.PrimaryKey()
	//353 vm_hosts
	tableVMHosts := VMHosts{}
	modelMap["vm_hosts"] = tableVMHosts
	primaryKeyMap["vm_hosts"] = tableVMHosts.PrimaryKey()
	//354 webshell_action
	tableWebshellAction := WebshellAction{}
	modelMap["webshell_action"] = tableWebshellAction
	primaryKeyMap["webshell_action"] = tableWebshellAction.PrimaryKey()
	//355 webshell_blacklist
	tableWebshellBlacklist := WebshellBlacklist{}
	modelMap["webshell_blacklist"] = tableWebshellBlacklist
	primaryKeyMap["webshell_blacklist"] = tableWebshellBlacklist.PrimaryKey()
	//356 webshell_capture
	tableWebshellCapture := WebshellCapture{}
	modelMap["webshell_capture"] = tableWebshellCapture
	primaryKeyMap["webshell_capture"] = tableWebshellCapture.PrimaryKey()
	//357 webshell_capture_dbcontent
	tableWebshellCaptureDbcontent := WebshellCaptureDbcontent{}
	modelMap["webshell_capture_dbcontent"] = tableWebshellCaptureDbcontent
	primaryKeyMap["webshell_capture_dbcontent"] = tableWebshellCaptureDbcontent.PrimaryKey()
	//358 webshell_group_permision
	tableWebshellGroupPermision := WebshellGroupPermision{}
	modelMap["webshell_group_permision"] = tableWebshellGroupPermision
	primaryKeyMap["webshell_group_permision"] = tableWebshellGroupPermision.PrimaryKey()
	//359 webshell_group_rel
	tableWebshellGroupRel := WebshellGroupRel{}
	modelMap["webshell_group_rel"] = tableWebshellGroupRel
	primaryKeyMap["webshell_group_rel"] = tableWebshellGroupRel.PrimaryKey()
	//360 webshell_role_rel
	tableWebshellRoleRel := WebshellRoleRel{}
	modelMap["webshell_role_rel"] = tableWebshellRoleRel
	primaryKeyMap["webshell_role_rel"] = tableWebshellRoleRel.PrimaryKey()
	//361 webshell_src_rel
	tableWebshellSrcRel := WebshellSrcRel{}
	modelMap["webshell_src_rel"] = tableWebshellSrcRel
	primaryKeyMap["webshell_src_rel"] = tableWebshellSrcRel.PrimaryKey()
	//362 webshell_watch
	tableWebshellWatch := WebshellWatch{}
	modelMap["webshell_watch"] = tableWebshellWatch
	primaryKeyMap["webshell_watch"] = tableWebshellWatch.PrimaryKey()
	//363 websocket_msg
	tableWebsocketMsg := WebsocketMsg{}
	modelMap["websocket_msg"] = tableWebsocketMsg
	primaryKeyMap["websocket_msg"] = tableWebsocketMsg.PrimaryKey()
	//364 websocket_msg_pid
	tableWebsocketMsgPid := WebsocketMsgPid{}
	modelMap["websocket_msg_pid"] = tableWebsocketMsgPid
	primaryKeyMap["websocket_msg_pid"] = tableWebsocketMsgPid.PrimaryKey()
	//365 websql_favorite
	tableWebsqlFavorite := WebsqlFavorite{}
	modelMap["websql_favorite"] = tableWebsqlFavorite
	primaryKeyMap["websql_favorite"] = tableWebsqlFavorite.PrimaryKey()
	//366 websql_favorite_group
	tableWebsqlFavoriteGroup := WebsqlFavoriteGroup{}
	modelMap["websql_favorite_group"] = tableWebsqlFavoriteGroup
	primaryKeyMap["websql_favorite_group"] = tableWebsqlFavoriteGroup.PrimaryKey()
	//367 websql_watch
	tableWebsqlWatch := WebsqlWatch{}
	modelMap["websql_watch"] = tableWebsqlWatch
	primaryKeyMap["websql_watch"] = tableWebsqlWatch.PrimaryKey()
	//368 webterminal_event
	tableWebterminalEvent := WebterminalEvent{}
	modelMap["webterminal_event"] = tableWebterminalEvent
	primaryKeyMap["webterminal_event"] = tableWebterminalEvent.PrimaryKey()
	//369 wechat_config
	tableWechatConfig := WechatConfig{}
	modelMap["wechat_config"] = tableWechatConfig
	primaryKeyMap["wechat_config"] = tableWechatConfig.PrimaryKey()
	//370 week_failure_report
	tableWeekFailureReport := WeekFailureReport{}
	modelMap["week_failure_report"] = tableWeekFailureReport
	primaryKeyMap["week_failure_report"] = tableWeekFailureReport.PrimaryKey()
	//371 week_issue_record
	tableWeekIssueRecord := WeekIssueRecord{}
	modelMap["week_issue_record"] = tableWeekIssueRecord
	primaryKeyMap["week_issue_record"] = tableWeekIssueRecord.PrimaryKey()
	//372 week_publish_req
	tableWeekPublishReq := WeekPublishReq{}
	modelMap["week_publish_req"] = tableWeekPublishReq
	primaryKeyMap["week_publish_req"] = tableWeekPublishReq.PrimaryKey()
	//373 week_user_team
	tableWeekUserTeam := WeekUserTeam{}
	modelMap["week_user_team"] = tableWeekUserTeam
	primaryKeyMap["week_user_team"] = tableWeekUserTeam.PrimaryKey()
	//374 work_station_duty
	tableWorkStationDuty := WorkStationDuty{}
	modelMap["work_station_duty"] = tableWorkStationDuty
	primaryKeyMap["work_station_duty"] = tableWorkStationDuty.PrimaryKey()
	//375 zabbix_maintenance_list
	tableZabbixMaintenanceList := ZabbixMaintenanceList{}
	modelMap["zabbix_maintenance_list"] = tableZabbixMaintenanceList
	primaryKeyMap["zabbix_maintenance_list"] = tableZabbixMaintenanceList.PrimaryKey()
	//376 zabbix_service_list
	tableZabbixServiceList := ZabbixServiceList{}
	modelMap["zabbix_service_list"] = tableZabbixServiceList
	primaryKeyMap["zabbix_service_list"] = tableZabbixServiceList.PrimaryKey()
}