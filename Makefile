NAME = my_git


all: $(NAME)

$(NAME):
	@go build -o $(NAME) main.go

fclean:
	@rm $(NAME)

re: fclean all

.PHONY:				all fclean re
