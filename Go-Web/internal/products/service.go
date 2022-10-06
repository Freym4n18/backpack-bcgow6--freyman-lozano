package products

type Service interface {
	GetAll() ([]Product, error)
	AddProduct(product Product) error
	UpdateProduct(id string, name string, price float64) (Product,error)
    DeleteProduct(id string) error
	Replace(id string, product Product) error
	GetOne(id string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
    return &service{
		repository: repository,
	}
}

func (s *service) GetOne(id string) (Product, error) {
    product, err := s.repository.GetById(id)
    return product, err
}

func(s *service) AddProduct(product Product) error {
	id, err := s.repository.FindNextId()
	if err!= nil {
        return err
    }
	product.Id = id
	return s.repository.AddProduct(product)
}

func(s *service) UpdateProduct(id string, name string, price float64) (Product,error) {
	product, err := s.repository.GetById(id)
    if err!= nil {
        return product,err
    }
	product.Name = name
    product.Price = price
	return s.repository.Update(id, name, price)
}

func(s *service) DeleteProduct(id string) error {
	return s.repository.Delete(id)
}

func(s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
    if err!= nil {
        return nil, err
    }
	return products, nil
}

func(s *service) Replace(id string, product Product) error {
	_, err := s.repository.GetById(id)
	if err!= nil {
        return err
    }
	product.Id = id
    return s.repository.Replace(id, product)
}