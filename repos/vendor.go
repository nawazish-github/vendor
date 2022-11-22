package repos

import models "github.com/nawazish-github/vendor/models"

type IVendorRepo interface {
	Write (v *models.Vendor)
}