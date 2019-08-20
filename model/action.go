/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package model

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/common"
	"github.com/unidoc/unipdf/v3/core"
)

type PdfActionType string

// // (Section 12.6.4 p. 417).
// See Table 198 - Action types
const (
	ActionTypeGoTo        PdfActionType = "GoTo"        // Go to a destination in the current document
	ActionTypeGoToR       PdfActionType = "GoToR"       // Go to remote, Go to a destination in another document
	ActionTypeGoToE       PdfActionType = "GoToE"       // Go to embedded, PDF 1.6, Got to a destination in an embedded file
	ActionTypeLaunch      PdfActionType = "Launch"      // Launch an application, usually to open a file
	ActionTypeThread      PdfActionType = "Thread"      // Begin reading an article thread
	ActionTypeURI         PdfActionType = "URI"         // Resolves a uniform resource identifier
	ActionTypeSound       PdfActionType = "Sound"       // Play a sound
	ActionTypeMovie       PdfActionType = "Movie"       // Play a movie
	ActionTypeHide        PdfActionType = "Hide"        // Set an annotation's Hidden flag
	ActionTypeNamed       PdfActionType = "Named"       // Execute an action predefined by the conforming reader
	ActionTypeSubmitForm  PdfActionType = "SubmitForm"  // Send data to a uniform resource locator
	ActionTypeResetForm   PdfActionType = "ResetForm"   // Set fields to their default values
	ActionTypeImportData  PdfActionType = "ImportData"  // Import field values from a file
	ActionTypeJavaScript  PdfActionType = "JavaScript"  // Execute a JavaScript script
	ActionTypeSetOCGState PdfActionType = "SetOCGState" // Set the states of optional content groups
	ActionTypeRendition   PdfActionType = "Rendition"   // Controls the playing of multimedia content
	ActionTypeTrans       PdfActionType = "Trans"       // Updates the display of a document, using a transition dictionary
	ActionTypeGoTo3DView  PdfActionType = "GoTo3DView"  // Set the current view of a 3D annotation
)

type PdfAction struct {
	context PdfModel

	Type core.PdfObject
	S    core.PdfObject
	Next core.PdfObject

	container *core.PdfIndirectObject
}

// GetContext returns the action context which contains the specific type-dependent context.
// The context represents the subaction.
func (a *PdfAction) GetContext() PdfModel {
	if a == nil {
		return nil
	}
	return a.context
}

// SetContext sets the sub action (context).
func (a *PdfAction) SetContext(ctx PdfModel) {
	a.context = ctx
}

// GetContainingPdfObject implements interface PdfModel.
func (a *PdfAction) GetContainingPdfObject() core.PdfObject {
	return a.container
}

// ToPdfObject implements interface PdfModel.
func (a *PdfAction) ToPdfObject() core.PdfObject {
	container := a.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.Clear()

	d.Set("Type", core.MakeName("Action"))
	d.SetIfNotNil("S", a.S)
	d.SetIfNotNil("Next", a.Next)

	return container
}

// String implements interface PdfObject.
func (a *PdfAction) String() string {
	s := ""

	obj, ok := a.ToPdfObject().(*core.PdfIndirectObject)
	if ok {
		s = fmt.Sprintf("%T: %s", a.context, obj.PdfObject.String())
	}

	return s
}

// PdfActionGoTo represents a GoTo action
type PdfActionGoTo struct {
	*PdfAction
	D core.PdfObject // name, byte string or array
}

// PdfActionGoToR represents a GoToR action
type PdfActionGoToR struct {
	*PdfAction
	F         core.PdfObject // file specification
	D         core.PdfObject // name, byte string or array
	NewWindow core.PdfObject
}

// PdfActionGoToE represents a GoToE action
type PdfActionGoToE struct {
	*PdfAction
	F         core.PdfObject // file specification
	D         core.PdfObject // name, byte string or array
	NewWindow core.PdfObject
	T         core.PdfObject
}

// PdfActionLaunch represents a launch action
type PdfActionLaunch struct {
	*PdfAction
	F         core.PdfObject // file specification
	Win       core.PdfObject
	Mac       core.PdfObject
	Unix      core.PdfObject
	NewWindow core.PdfObject
}

// PdfActionThread represents a thread action
type PdfActionThread struct {
	*PdfAction
	F core.PdfObject // file specification
	D core.PdfObject
	B core.PdfObject
}

// PdfActionURI represents an URI action
type PdfActionURI struct {
	*PdfAction
	URI   core.PdfObject
	IsMap core.PdfObject
}

// PdfActionSound represents a sound action
type PdfActionSound struct {
	*PdfAction
	Sound       core.PdfObject
	Volume      core.PdfObject
	Synchronous core.PdfObject
	Repeat      core.PdfObject
	Mix         core.PdfObject
}

// PdfActionMovie represents a movie action
type PdfActionMovie struct {
	*PdfAction
	Annotation core.PdfObject
	T          core.PdfObject
	Operation  core.PdfObject
}

// PdfActionHide represents a hide action
type PdfActionHide struct {
	*PdfAction
	T core.PdfObject
	H core.PdfObject
}

// PdfActionNamed represents a named action
type PdfActionNamed struct {
	*PdfAction
	N core.PdfObject
}

// PdfActionSubmitForm represents a submitForm action
type PdfActionSubmitForm struct {
	*PdfAction
	F core.PdfObject // File specification
	Fields core.PdfObject
	Flags core.PdfObject
}

// PdfActionResetForm represents a resetForm action
type PdfActionResetForm struct {
	*PdfAction
	Fields core.PdfObject
	Flags core.PdfObject
}

// PdfActionImportData represents a importData action
type PdfActionImportData struct {
	*PdfAction
	F core.PdfObject // File specification
}

// PdfActionSetOCGState represents a SetOCGState action
type PdfActionSetOCGState struct {
	*PdfAction
	State      core.PdfObject
	PreserveRB core.PdfObject
}

// PdfActionRendition represents a Rendition action
type PdfActionRendition struct {
	*PdfAction
	R  core.PdfObject
	AN core.PdfObject
	OP core.PdfObject
	JS core.PdfObject
}

// PdfActionTrans represents a trans action
type PdfActionTrans struct {
	*PdfAction
	Trans core.PdfObject
}

// PdfActionGoTo3DView represents a GoTo3DView action
type PdfActionGoTo3DView struct {
	*PdfAction
	TA core.PdfObject
	V  core.PdfObject
}

// PdfActionJavascript represents a javaScript action
type PdfActionJavaScript struct {
	*PdfAction
	JS core.PdfObject
}

// NewPdfAction returns an initialized generic PDF action model.
func NewPdfAction() *PdfAction {
	action := &PdfAction{}
	action.container = core.MakeIndirectObject(core.MakeDict())
	return action
}

// NewPdfActionGoTo returns a new "go to" action
func NewPdfActionGoTo() *PdfActionGoTo {
	action := NewPdfAction()
	goToAction := &PdfActionGoTo{}
	goToAction.PdfAction = action
	action.SetContext(goToAction)
	return goToAction
}

// NewPdfActionGoToR returns a new "go to remote" action
func NewPdfActionGoToR() *PdfActionGoToR {
	action := NewPdfAction()
	goToRAction := &PdfActionGoToR{}
	goToRAction.PdfAction = action
	action.SetContext(goToRAction)
	return goToRAction
}

// NewPdfActionGoToE returns a new "go to embedded" action
func NewPdfActionGoToE() *PdfActionGoToE {
	action := NewPdfAction()
	goToEAction := &PdfActionGoToE{}
	goToEAction.PdfAction = action
	action.SetContext(goToEAction)
	return goToEAction
}

// NewPdfActionLaunch returns a new "launch" action
func NewPdfActionLaunch() *PdfActionLaunch {
	action := NewPdfAction()
	launchAction := &PdfActionLaunch{}
	launchAction.PdfAction = action
	action.SetContext(launchAction)
	return launchAction
}

// NewPdfActionThread returns a new "thread" action
func NewPdfActionThread() *PdfActionThread {
	action := NewPdfAction()
	threadAction := &PdfActionThread{}
	threadAction.PdfAction = action
	action.SetContext(threadAction)
	return threadAction
}

// NewPdfActionURI returns a new "Uri" action
func NewPdfActionURI() *PdfActionURI {
	action := NewPdfAction()
	uriAction := &PdfActionURI{}
	uriAction.PdfAction = action
	action.SetContext(uriAction)
	return uriAction
}

// NewPdfActionSound returns a new "sound" action
func NewPdfActionSound() *PdfActionSound {
	action := NewPdfAction()
	soundAction := &PdfActionSound{}
	soundAction.PdfAction = action
	action.SetContext(soundAction)
	return soundAction
}

// NewPdfActionMovie returns a new "movie" action
func NewPdfActionMovie() *PdfActionMovie {
	action := NewPdfAction()
	movieAction := &PdfActionMovie{}
	movieAction.PdfAction = action
	action.SetContext(movieAction)
	return movieAction
}

// NewPdfActionHide returns a new "hide" action
func NewPdfActionHide() *PdfActionHide {
	action := NewPdfAction()
	hideAction := &PdfActionHide{}
	hideAction.PdfAction = action
	action.SetContext(hideAction)
	return hideAction
}

// NewPdfActionNamed returns a new "named" action
func NewPdfActionNamed() *PdfActionNamed {
	action := NewPdfAction()
	namedAction := &PdfActionNamed{}
	namedAction.PdfAction = action
	action.SetContext(namedAction)
	return namedAction
}

// NewPdfActionSubmitForm returns a new "submit form" action
func NewPdfActionSubmitForm() *PdfActionSubmitForm {
	action := NewPdfAction()
	submitFormAction := &PdfActionSubmitForm{}
	submitFormAction.PdfAction = action
	action.SetContext(submitFormAction)
	return submitFormAction
}

// NewPdfActionResetForm returns a new "reset form" action
func NewPdfActionResetForm() *PdfActionResetForm {
	action := NewPdfAction()
	resetFormAction := &PdfActionResetForm{}
	resetFormAction.PdfAction = action
	action.SetContext(resetFormAction)
	return resetFormAction
}

// NewPdfActionImportData returns a new "import data" action
func NewPdfActionImportData() *PdfActionImportData {
	action := NewPdfAction()
	importDataAction := &PdfActionImportData{}
	importDataAction.PdfAction = action
	action.SetContext(importDataAction)
	return importDataAction
}

// NewPdfActionSetOCGState returns a new "named" action
func NewPdfActionSetOCGState() *PdfActionSetOCGState {
	action := NewPdfAction()
	setOCGStateAction := &PdfActionSetOCGState{}
	setOCGStateAction.PdfAction = action
	action.SetContext(setOCGStateAction)
	return setOCGStateAction
}

// NewPdfActionRendition returns a new "rendition" action
func NewPdfActionRendition() *PdfActionRendition {
	action := NewPdfAction()
	renditionAction := &PdfActionRendition{}
	renditionAction.PdfAction = action
	action.SetContext(renditionAction)
	return renditionAction
}

// NewPdfActionTrans returns a new "trans" action
func NewPdfActionTrans() *PdfActionTrans {
	action := NewPdfAction()
	transAction := &PdfActionTrans{}
	transAction.PdfAction = action
	action.SetContext(transAction)
	return transAction
}

// NewPdfActionGoTo3DView returns a new "goTo3DView" action
func NewPdfActionGoTo3DView() *PdfActionGoTo3DView {
	action := NewPdfAction()
	goTo3DViewAction := &PdfActionGoTo3DView{}
	goTo3DViewAction.PdfAction = action
	action.SetContext(goTo3DViewAction)
	return goTo3DViewAction
}

// NewPdfActionJavaScript returns a new "javaScript" action
func NewPdfActionJavaScript() *PdfActionJavaScript {
	action := NewPdfAction()
	javaScriptAction := &PdfActionJavaScript{}
	javaScriptAction.PdfAction = action
	action.SetContext(javaScriptAction)
	return javaScriptAction
}

// ToPdfObject implements interface PdfModel.
func (gotoAct *PdfActionGoTo) ToPdfObject() core.PdfObject {
	gotoAct.PdfAction.ToPdfObject()
	container := gotoAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeGoTo)))
	d.SetIfNotNil("D", gotoAct.D)
	return container
}

// ToPdfObject implements interface PdfModel.
func (gotoRAct *PdfActionGoToR) ToPdfObject() core.PdfObject {
	gotoRAct.PdfAction.ToPdfObject()
	container := gotoRAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeGoToR)))
	d.SetIfNotNil("F", gotoRAct.F)
	d.SetIfNotNil("D", gotoRAct.D)
	d.SetIfNotNil("NewWindow", gotoRAct.NewWindow)
	return container
}

// ToPdfObject implements interface PdfModel.
func (gotoEAct *PdfActionGoToE) ToPdfObject() core.PdfObject {
	gotoEAct.PdfAction.ToPdfObject()
	container := gotoEAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeGoToE)))
	d.SetIfNotNil("F", gotoEAct.F)
	d.SetIfNotNil("D", gotoEAct.D)
	d.SetIfNotNil("NewWindow", gotoEAct.NewWindow)
	d.SetIfNotNil("T", gotoEAct.T)
	return container
}

// ToPdfObject implements interface PdfModel.
func (launchAct *PdfActionLaunch) ToPdfObject() core.PdfObject {
	launchAct.PdfAction.ToPdfObject()
	container := launchAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeLaunch)))
	d.SetIfNotNil("F", launchAct.F)
	d.SetIfNotNil("Win", launchAct.Win)
	d.SetIfNotNil("Mac", launchAct.Mac)
	d.SetIfNotNil("Unix", launchAct.Unix)
	d.SetIfNotNil("NewWindow", launchAct.NewWindow)
	return container
}

// ToPdfObject implements interface PdfModel.
func (threadAct *PdfActionThread) ToPdfObject() core.PdfObject {
	threadAct.PdfAction.ToPdfObject()
	container := threadAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeThread)))
	d.SetIfNotNil("F", threadAct.F)
	d.SetIfNotNil("D", threadAct.D)
	d.SetIfNotNil("B", threadAct.B)
	return container
}

// ToPdfObject implements interface PdfModel.
func (uriAct *PdfActionURI) ToPdfObject() core.PdfObject {
	uriAct.PdfAction.ToPdfObject()
	container := uriAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeURI)))
	d.SetIfNotNil("URI", uriAct.URI)
	d.SetIfNotNil("IsMap", uriAct.IsMap)
	return container
}

// ToPdfObject implements interface PdfModel.
func (soundAct *PdfActionSound) ToPdfObject() core.PdfObject {
	soundAct.PdfAction.ToPdfObject()
	container := soundAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeSound)))
	d.SetIfNotNil("Sound", soundAct.Sound)
	d.SetIfNotNil("Volume", soundAct.Volume)
	d.SetIfNotNil("Synchronous", soundAct.Synchronous)
	d.SetIfNotNil("Repeat", soundAct.Repeat)
	d.SetIfNotNil("Mix", soundAct.Mix)
	return container
}

// ToPdfObject implements interface PdfModel.
func (movieAct *PdfActionMovie) ToPdfObject() core.PdfObject {
	movieAct.PdfAction.ToPdfObject()
	container := movieAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeMovie)))
	d.SetIfNotNil("Annotation", movieAct.Annotation)
	d.SetIfNotNil("T", movieAct.T)
	d.SetIfNotNil("Operation", movieAct.Operation)
	return container
}

// ToPdfObject implements interface PdfModel.
func (hideAct *PdfActionHide) ToPdfObject() core.PdfObject {
	hideAct.PdfAction.ToPdfObject()
	container := hideAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeHide)))
	d.SetIfNotNil("T", hideAct.T)
	d.SetIfNotNil("H", hideAct.H)
	return container
}

// ToPdfObject implements interface PdfModel.
func (namedAct *PdfActionNamed) ToPdfObject() core.PdfObject {
	namedAct.PdfAction.ToPdfObject()
	container := namedAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeNamed)))
	d.SetIfNotNil("N", namedAct.N)
	return container
}

// ToPdfObject implements interface PdfModel.
func (submitFormAct *PdfActionSubmitForm) ToPdfObject() core.PdfObject {
	submitFormAct.PdfAction.ToPdfObject()
	container := submitFormAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeSubmitForm)))
	d.SetIfNotNil("F", submitFormAct.F)
	d.SetIfNotNil("Fields", submitFormAct.Fields)
	d.SetIfNotNil("Flags", submitFormAct.Flags)
	return container
}

// ToPdfObject implements interface PdfModel.
func (resetFormAct *PdfActionResetForm) ToPdfObject() core.PdfObject {
	resetFormAct.PdfAction.ToPdfObject()
	container := resetFormAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeResetForm)))
	d.SetIfNotNil("Fields", resetFormAct.Fields)
	d.SetIfNotNil("Flags", resetFormAct.Flags)
	return container
}

// ToPdfObject implements interface PdfModel.
func (importDataAct *PdfActionImportData) ToPdfObject() core.PdfObject {
	importDataAct.PdfAction.ToPdfObject()
	container := importDataAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeImportData)))
	d.SetIfNotNil("F", importDataAct.F)
	return container
}

// ToPdfObject implements interface PdfModel.
func (setOCGStateAct *PdfActionSetOCGState) ToPdfObject() core.PdfObject {
	setOCGStateAct.PdfAction.ToPdfObject()
	container := setOCGStateAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeSetOCGState)))
	d.SetIfNotNil("State", setOCGStateAct.State)
	d.SetIfNotNil("PreserveRB", setOCGStateAct.PreserveRB)
	return container
}

// ToPdfObject implements interface PdfModel.
func (renditionAct *PdfActionRendition) ToPdfObject() core.PdfObject {
	renditionAct.PdfAction.ToPdfObject()
	container := renditionAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeRendition)))
	d.SetIfNotNil("R", renditionAct.R)
	d.SetIfNotNil("AN", renditionAct.AN)
	d.SetIfNotNil("OP", renditionAct.OP)
	d.SetIfNotNil("JS", renditionAct.JS)
	return container
}

// ToPdfObject implements interface PdfModel.
func (transAct *PdfActionTrans) ToPdfObject() core.PdfObject {
	transAct.PdfAction.ToPdfObject()
	container := transAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeTrans)))
	d.SetIfNotNil("Trans", transAct.Trans)
	return container
}

// ToPdfObject implements interface PdfModel.
func (goTo3DViewAct *PdfActionGoTo3DView) ToPdfObject() core.PdfObject {
	goTo3DViewAct.PdfAction.ToPdfObject()
	container := goTo3DViewAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeGoTo3DView)))
	d.SetIfNotNil("TA", goTo3DViewAct.TA)
	d.SetIfNotNil("V", goTo3DViewAct.V)
	return container
}

// ToPdfObject implements interface PdfModel.
func (javaScriptAct *PdfActionJavaScript) ToPdfObject() core.PdfObject {
	javaScriptAct.PdfAction.ToPdfObject()
	container := javaScriptAct.container
	d := container.PdfObject.(*core.PdfObjectDictionary)

	d.SetIfNotNil("S", core.MakeName(string(ActionTypeJavaScript)))
	d.SetIfNotNil("JS", javaScriptAct.JS)
	return container
}

// Used for PDF parsing.  Loads a PDF action model from a PDF dictionary.
// Loads the common PDF action dictionary, and anything needed for the action subtype.
func (r *PdfReader) newPdfActionFromIndirectObject(container *core.PdfIndirectObject) (*PdfAction, error) {
	d, isDict := container.PdfObject.(*core.PdfObjectDictionary)
	if !isDict {
		return nil, fmt.Errorf("action indirect object not containing a dictionary")
	}

	// Check if cached, return cached model if exists.
	if model := r.modelManager.GetModelFromPrimitive(d); model != nil {
		action, ok := model.(*PdfAction)
		if !ok {
			return nil, fmt.Errorf("cached model not a PDF action")
		}
		return action, nil
	}

	action := &PdfAction{}
	action.container = container
	r.modelManager.Register(d, action)

	if obj := d.Get("Type"); obj != nil {
		str, ok := obj.(*core.PdfObjectName)
		if !ok {
			common.Log.Trace("Incompatibility! Invalid type of Type (%T) - should be Name", obj)
		} else {
			if *str != "Action" {
				// Log a debug message.
				// Not returning an error on this.
				common.Log.Trace("Unsuspected Type != Action (%s)", *str)
			}
		}
	}

	if obj := d.Get("Next"); obj != nil {
		action.Next = obj
	}

	if obj := d.Get("S"); obj != nil {
		action.S = obj
	}

	actionType, ok := action.S.(*core.PdfObjectName)
	if !ok {
		common.Log.Debug("ERROR: Invalid S object type != name (%T)", action.S)
		return nil, fmt.Errorf("invalid S object type != name (%T)", action.S)
	}
	switch string(*actionType) {
	case string(ActionTypeGoTo):
		ctx, err := r.newPdfActionGotoFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeGoToR):
		ctx, err := r.newPdfActionGotoRFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeGoToE):
		ctx, err := r.newPdfActionGotoEFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeLaunch):
		ctx, err := r.newPdfActionLaunchFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeThread):
		ctx, err := r.newPdfActionThreadFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeURI):
		ctx, err := r.newPdfActionURIFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeSound):
		ctx, err := r.newPdfActionSoundFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeMovie):
		ctx, err := r.newPdfActionMovieFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeHide):
		ctx, err := r.newPdfActionHideFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeNamed):
		ctx, err := r.newPdfActionNamedFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeSubmitForm):
		ctx, err := r.newPdfActionSubmitFormFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeResetForm):
		ctx, err := r.newPdfActionResetFormFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeImportData):
		ctx, err := r.newPdfActionImportDataFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeSetOCGState):
		ctx, err := r.newPdfActionSetOCGStateFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeRendition):
		ctx, err := r.newPdfActionRenditionFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeTrans):
		ctx, err := r.newPdfActionTransFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeGoTo3DView):
		ctx, err := r.newPdfActionGoTo3DViewFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil
	case string(ActionTypeJavaScript):
		ctx, err := r.newPdfActionJavaScriptFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAction = action
		action.context = ctx
		return action, nil

	}

	common.Log.Debug("ERROR: Ignoring unknown action: %s", *actionType)
	return nil, nil
}

func (r *PdfReader) newPdfActionGotoFromDict(d *core.PdfObjectDictionary) (*PdfActionGoTo, error) {
	action := PdfActionGoTo{}

	action.D = d.Get("D")

	return &action, nil
}

func (r *PdfReader) newPdfActionGotoRFromDict(d *core.PdfObjectDictionary) (*PdfActionGoToR, error) {
	action := PdfActionGoToR{}

	action.F = d.Get("F")
	action.D = d.Get("D")
	action.NewWindow = d.Get("NewWindow")

	return &action, nil
}

func (r *PdfReader) newPdfActionGotoEFromDict(d *core.PdfObjectDictionary) (*PdfActionGoToE, error) {
	action := PdfActionGoToE{}

	action.F = d.Get("F")
	action.D = d.Get("D")
	action.NewWindow = d.Get("NewWindow")
	action.T = d.Get("T")

	return &action, nil
}

func (r *PdfReader) newPdfActionLaunchFromDict(d *core.PdfObjectDictionary) (*PdfActionLaunch, error) {
	action := PdfActionLaunch{}

	action.F = d.Get("F")
	action.Win = d.Get("Win")
	action.Mac = d.Get("Mac")
	action.Unix = d.Get("Unix")
	action.NewWindow = d.Get("NewWindow")

	return &action, nil
}

func (r *PdfReader) newPdfActionThreadFromDict(d *core.PdfObjectDictionary) (*PdfActionThread, error) {
	action := PdfActionThread{}

	action.F = d.Get("F")
	action.D = d.Get("D")
	action.B = d.Get("B")

	return &action, nil
}

func (r *PdfReader) newPdfActionURIFromDict(d *core.PdfObjectDictionary) (*PdfActionURI, error) {
	action := PdfActionURI{}

	action.URI = d.Get("URI")
	action.IsMap = d.Get("IsMap")

	return &action, nil
}

func (r *PdfReader) newPdfActionSoundFromDict(d *core.PdfObjectDictionary) (*PdfActionSound, error) {
	action := PdfActionSound{}

	action.Sound = d.Get("Sound")
	action.Volume = d.Get("Volume")
	action.Synchronous = d.Get("Synchronous")
	action.Repeat = d.Get("Repeat")
	action.Mix = d.Get("Mix")

	return &action, nil
}

func (r *PdfReader) newPdfActionMovieFromDict(d *core.PdfObjectDictionary) (*PdfActionMovie, error) {
	action := PdfActionMovie{}

	action.Annotation = d.Get("Annotation")
	action.T = d.Get("T")
	action.Operation = d.Get("Operation")

	return &action, nil
}

func (r *PdfReader) newPdfActionHideFromDict(d *core.PdfObjectDictionary) (*PdfActionHide, error) {
	action := PdfActionHide{}

	action.T = d.Get("T")
	action.H = d.Get("H")

	return &action, nil
}

func (r *PdfReader) newPdfActionNamedFromDict(d *core.PdfObjectDictionary) (*PdfActionNamed, error) {
	action := PdfActionNamed{}

	action.N = d.Get("N")

	return &action, nil
}

func (r *PdfReader) newPdfActionSubmitFormFromDict(d *core.PdfObjectDictionary) (*PdfActionSubmitForm, error) {
	action := PdfActionSubmitForm{}

	action.F = d.Get("F")
	action.Fields = d.Get("Fields")
	action.Flags = d.Get("Flags")

	return &action, nil
}

func (r *PdfReader) newPdfActionResetFormFromDict(d *core.PdfObjectDictionary) (*PdfActionResetForm, error) {
	action := PdfActionResetForm{}

	action.Fields = d.Get("Fields")
	action.Flags = d.Get("Flags")

	return &action, nil
}

func (r *PdfReader) newPdfActionImportDataFromDict(d *core.PdfObjectDictionary) (*PdfActionImportData, error) {
	action := PdfActionImportData{}

	action.F = d.Get("F")

	return &action, nil
}

func (r *PdfReader) newPdfActionSetOCGStateFromDict(d *core.PdfObjectDictionary) (*PdfActionSetOCGState, error) {
	action := PdfActionSetOCGState{}

	action.State = d.Get("State")
	action.PreserveRB = d.Get("PreserveRB")

	return &action, nil
}

func (r *PdfReader) newPdfActionRenditionFromDict(d *core.PdfObjectDictionary) (*PdfActionRendition, error) {
	action := PdfActionRendition{}

	action.R = d.Get("R")
	action.AN = d.Get("AN")
	action.OP = d.Get("OP")
	action.JS = d.Get("JS")

	return &action, nil
}

func (r *PdfReader) newPdfActionTransFromDict(d *core.PdfObjectDictionary) (*PdfActionTrans, error) {
	action := PdfActionTrans{}

	action.Trans = d.Get("Trans")

	return &action, nil
}

func (r *PdfReader) newPdfActionGoTo3DViewFromDict(d *core.PdfObjectDictionary) (*PdfActionGoTo3DView, error) {
	action := PdfActionGoTo3DView{}

	action.TA = d.Get("TA")
	action.V = d.Get("V")

	return &action, nil
}

func (r *PdfReader) newPdfActionJavaScriptFromDict(d *core.PdfObjectDictionary) (*PdfActionJavaScript, error) {
	action := PdfActionJavaScript{}

	action.JS = d.Get("JS")

	return &action, nil
}
