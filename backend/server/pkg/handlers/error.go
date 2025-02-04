package handlers

//func ErrorHandler(c *gin.Context, currentError error) {
//	switch err := currentError.(type) {
//
//	case *enum.ValidationError:
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Validation Error": err.Error()})
//
//	case *e.DatabaseError:
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Database Error": err.Error()})
//
//	case *e.ServerError:
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Server Error": err.Error()})
//
//	default:
//		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"Unknown Error": err.Error()})
//	}
//}
