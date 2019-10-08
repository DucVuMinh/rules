package mem

import (
	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/rete/internal/types"
)

//Holds a tuple reference and related state

type reteHandleImpl struct {
	types.NwElemIdImpl
	tuple    model.Tuple
	tupleKey model.TupleKey
	status   types.ReteHandleStatus
}

func newReteHandleImpl(nw types.Network, tuple model.Tuple, status types.ReteHandleStatus) types.ReteHandle {
	h1 := reteHandleImpl{}
	h1.initHandleImpl(nw, tuple, status)
	return &h1
}

func (hdl *reteHandleImpl) SetTuple(tuple model.Tuple) {
	hdl.tuple = tuple
	hdl.tupleKey = tuple.GetKey()
}

func (hdl *reteHandleImpl) initHandleImpl(nw types.Network, tuple model.Tuple, status types.ReteHandleStatus) {
	hdl.SetID(nw)
	hdl.SetTuple(tuple)
	hdl.tupleKey = tuple.GetKey()
	hdl.status = status
}

func (hdl *reteHandleImpl) GetTuple() model.Tuple {
	return hdl.tuple
}

func (hdl *reteHandleImpl) GetTupleKey() model.TupleKey {
	return hdl.tupleKey
}

func (hdl *reteHandleImpl) SetStatus(status types.ReteHandleStatus) {
	hdl.status = status
}

func (hdl *reteHandleImpl) GetStatus() types.ReteHandleStatus {
	return hdl.status
}

func (hdl *reteHandleImpl) AddJoinTableRowRef(joinTableRowVar types.JoinTableRow, joinTableVar types.JoinTable) {
	hdl.Nw.GetJtRefService().AddEntry(hdl, joinTableVar.GetName(), joinTableRowVar.GetID())
}