package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book UpdateBookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, updateBookRequest UpdateBookRequest) (Book, error) {
	book, err := s.repository.FindByID(ID)

	price, _ := updateBookRequest.Price.Int64()
	rating, _ := updateBookRequest.Rating.Int64()
	discount, _ := updateBookRequest.Discount.Int64()

	book.Title = updateBookRequest.Title
	book.Price = int(price)
	book.Description = updateBookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	newBook, err := s.repository.Delete(book)
	return newBook, err
}
