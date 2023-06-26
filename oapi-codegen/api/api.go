package api

import (
	"context"
	"fmt"
	"net/http"
)

type MyStrictServer struct {
	Pets   map[int64]Pet
	NextId int64
}

var _ StrictServerInterface = (*MyStrictServer)(nil)

func NewMyStrictServer() *MyStrictServer {
	return &MyStrictServer{
		Pets:   make(map[int64]Pet),
		NextId: 1000,
	}
}

func (ss *MyStrictServer) FindPets(ctx context.Context, request FindPetsRequestObject) (FindPetsResponseObject, error) {
	var result []Pet = []Pet{}

	for _, pet := range ss.Pets {
		if request.Params.Tags != nil {
			for _, t := range *request.Params.Tags {
				if pet.Tag != nil && (*pet.Tag == t) {
					result = append(result, pet)
				}
			}
		} else {
			result = append(result, pet)
		}

		if request.Params.Limit != nil {
			l := int(*request.Params.Limit)
			if len(result) >= l {
				break
			}
		}
	}

	return FindPets200JSONResponse(result), nil
}

func (ss *MyStrictServer) AddPet(ctx context.Context, request AddPetRequestObject) (AddPetResponseObject, error) {
	var pet Pet
	pet.Name = request.Body.Name
	pet.Tag = request.Body.Tag
	pet.Id = ss.NextId
	ss.NextId = ss.NextId + 1

	ss.Pets[pet.Id] = pet

	return AddPet200JSONResponse(pet), nil
}

func (ss *MyStrictServer) FindPetByID(ctx context.Context, request FindPetByIDRequestObject) (FindPetByIDResponseObject, error) {
	pet, found := ss.Pets[request.Id]
	if !found {
		return FindPetByIDdefaultJSONResponse{StatusCode: http.StatusNotFound, Body: Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Could not find pet with ID %d", request.Id)}}, nil
	}

	return FindPetByID200JSONResponse(pet), nil
}

func (ss *MyStrictServer) DeletePet(ctx context.Context, request DeletePetRequestObject) (DeletePetResponseObject, error) {
	_, found := ss.Pets[request.Id]
	if !found {
		return DeletePetdefaultJSONResponse{StatusCode: http.StatusNotFound, Body: Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Could not find pet with ID %d", request.Id)}}, nil
	}
	delete(ss.Pets, request.Id)

	return DeletePet204Response{}, nil
}
