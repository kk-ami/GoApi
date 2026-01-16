package middleware 

import (
	"errors"
	"net/http"

	"github.com/kk-ami/GoApi/api"
	"github.com/kk-ami/GoApi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

