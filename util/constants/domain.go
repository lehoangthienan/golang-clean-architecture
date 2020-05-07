package constants

// ProductType type
type ProductType string

func (t ProductType) String() string {
	return string(t)
}

// RequestType type
type RequestType string

// StatusType type
type StatusType string

const (
	// ServiceProductType type
	ServiceProductType ProductType = "service"
)

// ProductStatus status
type ProductStatus string

func (s ProductStatus) String() string {
	return string(s)
}

const (
	// VisibleProductStatus status
	VisibleProductStatus ProductStatus = "visible"
	// InvisibleProductStatus status
	InvisibleProductStatus ProductStatus = "invisible"
)

// ClinicRequestStatus status
type ClinicRequestStatus string

func (s ClinicRequestStatus) String() string {
	return string(s)
}

const (
	ClinicRequestAcceptStatus  ClinicRequestStatus = "accept"
	ClinicRequestRejectStatus  ClinicRequestStatus = "reject"
	ClinicRequestPendingStatus ClinicRequestStatus = "pending"
)

type ClinicRequestType string

func (s ClinicRequestType) String() string {
	return string(s)
}

const (
	ClinicRequestRequestType ClinicRequestType = "request"
)

// OrderStatus status
type OrderStatus string

func (s OrderStatus) String() string {
	return string(s)
}

const (
	OrderDoneStatus    OrderStatus = "done"
	OrderPendingStatus OrderStatus = "pending"
)

// UserRole status
type UserRole string

func (s UserRole) String() string {
	return string(s)
}

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleSale  UserRole = "sale"
	UserRoleUser  UserRole = "user"
)
