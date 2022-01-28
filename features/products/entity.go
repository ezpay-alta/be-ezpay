package products

type Core struct {
	ID      int
	Name    string
	TypeID  int
	Type    TypeCore
	Nominal int
	Price   int
}

type TypeCore struct {
	ID          int
	Name        string
	Description string
}

type Business interface {
	CreateType(productType TypeCore) error
	GetAllTypes() ([]TypeCore, error)
	GetTypeById(typeId int) (TypeCore, error)
	UpdateTypeById(typeId int, data TypeCore) error
	DeleteTypeById(typeId int) error

	CreateProduct(data Core) error
	GetAllProducts() ([]Core, error)
	GetProductByName(name string) (int, error)
	GetAllProductsByType(typeProduct string) ([]Core, error)
	GetProductById(productId int) (Core, error)
	UpdateProductById(productId int, data Core) error
	DeleteProductById(productId int) error
}

type Data interface {
	CreateType(productType TypeCore) error
	GetAllTypes() ([]TypeCore, error)
	GetTypeById(typeId int) (TypeCore, error)
	UpdateTypeById(typeId int, data TypeCore) error
	DeleteTypeById(typeId int) error

	CreateProduct(data Core) error
	GetAllProducts() ([]Core, error)
	GetProductByName(name string) (int, error)
	GetAllProductsByType(typeProduct string) ([]Core, error)
	GetProductById(productId int) (Core, error)
	UpdateProductById(productId int, data Core) error
	DeleteProductById(productId int) error
}
