package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
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
	newBook := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}
	book, err := s.repository.Create(newBook)

	return book, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {

	book, err := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = book.Description
	book.Rating = book.Rating

	newBook, err := s.repository.Update(book)

	return newBook, err
}
