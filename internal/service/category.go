package service




type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category,
}

