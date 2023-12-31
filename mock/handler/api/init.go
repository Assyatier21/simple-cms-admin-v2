// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/handler/api/init.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockDeliveryHandler is a mock of DeliveryHandler interface.
type MockDeliveryHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDeliveryHandlerMockRecorder
}

// MockDeliveryHandlerMockRecorder is the mock recorder for MockDeliveryHandler.
type MockDeliveryHandlerMockRecorder struct {
	mock *MockDeliveryHandler
}

// NewMockDeliveryHandler creates a new mock instance.
func NewMockDeliveryHandler(ctrl *gomock.Controller) *MockDeliveryHandler {
	mock := &MockDeliveryHandler{ctrl: ctrl}
	mock.recorder = &MockDeliveryHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeliveryHandler) EXPECT() *MockDeliveryHandlerMockRecorder {
	return m.recorder
}

// DeleteArticle mocks base method.
func (m *MockDeliveryHandler) DeleteArticle(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockDeliveryHandlerMockRecorder) DeleteArticle(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockDeliveryHandler)(nil).DeleteArticle), c)
}

// DeleteCategory mocks base method.
func (m *MockDeliveryHandler) DeleteCategory(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockDeliveryHandlerMockRecorder) DeleteCategory(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockDeliveryHandler)(nil).DeleteCategory), c)
}

// GetArticleDetails mocks base method.
func (m *MockDeliveryHandler) GetArticleDetails(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleDetails", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticleDetails indicates an expected call of GetArticleDetails.
func (mr *MockDeliveryHandlerMockRecorder) GetArticleDetails(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleDetails", reflect.TypeOf((*MockDeliveryHandler)(nil).GetArticleDetails), c)
}

// GetArticles mocks base method.
func (m *MockDeliveryHandler) GetArticles(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockDeliveryHandlerMockRecorder) GetArticles(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockDeliveryHandler)(nil).GetArticles), c)
}

// GetCategories mocks base method.
func (m *MockDeliveryHandler) GetCategories(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockDeliveryHandlerMockRecorder) GetCategories(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockDeliveryHandler)(nil).GetCategories), c)
}

// GetCategoryDetails mocks base method.
func (m *MockDeliveryHandler) GetCategoryDetails(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryDetails", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetCategoryDetails indicates an expected call of GetCategoryDetails.
func (mr *MockDeliveryHandlerMockRecorder) GetCategoryDetails(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryDetails", reflect.TypeOf((*MockDeliveryHandler)(nil).GetCategoryDetails), c)
}

// InsertArticle mocks base method.
func (m *MockDeliveryHandler) InsertArticle(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertArticle", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertArticle indicates an expected call of InsertArticle.
func (mr *MockDeliveryHandlerMockRecorder) InsertArticle(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertArticle", reflect.TypeOf((*MockDeliveryHandler)(nil).InsertArticle), c)
}

// InsertCategory mocks base method.
func (m *MockDeliveryHandler) InsertCategory(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCategory", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCategory indicates an expected call of InsertCategory.
func (mr *MockDeliveryHandlerMockRecorder) InsertCategory(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCategory", reflect.TypeOf((*MockDeliveryHandler)(nil).InsertCategory), c)
}

// UpdateArticle mocks base method.
func (m *MockDeliveryHandler) UpdateArticle(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticle", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateArticle indicates an expected call of UpdateArticle.
func (mr *MockDeliveryHandlerMockRecorder) UpdateArticle(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockDeliveryHandler)(nil).UpdateArticle), c)
}

// UpdateCategory mocks base method.
func (m *MockDeliveryHandler) UpdateCategory(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockDeliveryHandlerMockRecorder) UpdateCategory(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockDeliveryHandler)(nil).UpdateCategory), c)
}
