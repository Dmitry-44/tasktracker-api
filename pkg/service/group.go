package service

import (
	"errors"
	"fmt"
	"strconv"
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"

	"github.com/gin-gonic/gin"
)

type GroupsService struct {
	groupsRepo repository.Groups
}

func NewGroupService(groupsRepo repository.Groups) *GroupsService {
	return &GroupsService{groupsRepo: groupsRepo}
}

func (s *GroupsService) GetAll(user int) (models.GroupList, error) {
	return s.groupsRepo.GetAll(user)
}

func (s *GroupsService) GetGroupById(ctx *gin.Context, user int) (models.Group, error) {
	group := models.Group{}
	groupId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return group, fmt.Errorf("server error: %v", err.Error())
	}
	ok := s.isUserBelongsToGroup(user, groupId)
	if !ok {
		return group, errors.New("server error")
	}
	group, err = s.groupsRepo.GetGroupById(groupId)
	if err != nil {
		return group, fmt.Errorf("server error: %v", err.Error())
	}
	return group, nil
}

func (s *GroupsService) CreateGroup(user int, group models.GroupData) (int, error) {
	createdGroupId, err := s.groupsRepo.CreateGroup(user, group)
	if err != nil {
		return 0, err
	}
	return createdGroupId, nil
}

func (s *GroupsService) DeleteGroup(ctx *gin.Context, user int) error {
	groupId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return fmt.Errorf("server error: %v", err.Error())
	}
	ok := s.isUserBelongsToGroup(user, groupId)
	if !ok {
		return errors.New("server error")
	}
	err = s.groupsRepo.DeleteGroupById(groupId)
	if err != nil {
		return fmt.Errorf("server error: %v", err.Error())
	}
	return nil
}

func (s *GroupsService) isUserBelongsToGroup(user int, group int) bool {
	res := false
	userGroups, err := s.groupsRepo.GetAll(user)
	if err != nil {
		return res
	}
	for i := range userGroups.Groups {
		if userGroups.Groups[i].Id == group {
			res = true
			break
		}
	}
	return res
}
