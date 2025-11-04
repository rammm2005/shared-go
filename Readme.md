# shared-go 🧩
Use This shared for response, utils and custom hooks in Go in Microservice Arch
Shared utilities module for Go microservices — designed for consistent
error handling, response format, logging, environment management, and validation.

### Features
- 🧱 Standardized JSON response structure
- ⚠️ Unified custom error system
- 🪵 Lightweight logger (stdout/stderr)
- ⚙️ Env loader with fallback support
- ✅ Struct validator helper

### Example usage

```go
import (
    "github.com/rammm2005/shared-go/errors"
    "github.com/rammm2005/shared-go/response"
)

func (h *Handler) GetUser(c echo.Context) error {
    user, err := h.Service.FindUser(...)
    if err != nil {
        return response.JSON(c, nil, errors.NotFound("user not found"))
    }
    return response.JSON(c, user, nil)
}
