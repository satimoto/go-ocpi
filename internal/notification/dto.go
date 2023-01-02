package notification

type NotificationDto map[string]interface{}

func CreateTokenAuthorizeNotificationDto(authorizationId string) NotificationDto {
	response := map[string]interface{}{
		"type":            TOKEN_AUTHORIZE,
		"authorizationId": authorizationId,
	}

	return response
}
