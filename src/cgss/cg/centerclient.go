package cg

import (
	"encoding/json"
	"errors"
	"sync"

	"cgss/ipc"
)

var _ ipc.Server
