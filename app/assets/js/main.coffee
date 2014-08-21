$ = window.$ = window.jQuery = require 'jquery'
require 'bootstrap'

collapseSelector = '.js-vocanew-collapse'
contactModalSelector = '#js-vocanew-contact-modal'
contactFormSelector = '.js-vocanew-contact-form'
contactFormTextareaSelector = '.js-vocanew-contact-form textarea'
contactFormFieldsetSelector = '.js-vocanew-contact-form-fieldset'
contactFormLoadingImgSelector = '.js-vocanew-loading-img'
formErrorSelector = '.js-vocanew-form-error'
errorMessageSelector = '.js-vocanew-error-message'
alertSelector = '.js-vocanew-alert'

setupToggleSidebarCollapse = ->
  $(collapseSelector).on 'click touchstart', '.js-vocanew-collapse-btn', (ev) ->
    do ev.preventDefault
    $(collapseSelector).toggleClass 'vocanew-collapse-active'

setupContactModal = ->
  $(collapseSelector).on 'click touchstart', '.js-vocanew-toggle-contact-modal', (ev) ->
    do ev.preventDefault
    $(contactModalSelector).modal 'toggle'

  $(contactModalSelector).on 'click touchstart', '.js-vocanew-close', (ev) ->
    do ev.preventDefault
    $(contactModalSelector).modal 'hide'
  .on 'submit', contactFormSelector, (ev) ->
    do ev.preventDefault
    $form = $(contactFormSelector)
    data = do $form.serializeArray

    toggleSubmitting true

    $.ajax
      url: $form.attr 'action'
      type: 'POST',
      data: data
    .done(_onSubmitSuccess)
    .fail(_onSubmitFail)
    .always(-> toggleSubmitting false)
  .on 'hidden.bs.modal', ->
    toggleSubmitting false
  .on 'shown.bs.modal', ->
    $(contactFormTextareaSelector).trigger 'focus'

_onSubmitSuccess = (data) ->
  return _onSubmitFail data unless data.success

  dfd = do $.Deferred
  savedTop = $(alertSelector).css 'top'

  dfd.then ->
    dfd = do $.Deferred

    $(alertSelector).show().animate
      top: 0
    ,
      duration: 200,
      complete: dfd.resolve

    do dfd.promise
  .then ->
    dfd = do $.Deferred

    setTimeout ->
      $(alertSelector).animate
        top: savedTop
      ,
        duration: 200,
        complete: dfd.resolve
    , 3000

    do dfd.promise
  .then ->
    do $(alertSelector).hide
  $(contactFormSelector).trigger 'reset'
  $(contactModalSelector).on 'hidden.bs.modal', notify = (ev) ->
    $(ev.currentTarget).off 'hidden.bs.modal', notify
    do $(formErrorSelector).hide
    do dfd.resolve
  .modal 'hide'

_onSubmitFail = (data) ->
  result = data.responseJSON or data
  $errorMsg = $(errorMessageSelector)

  if result.error_messages
    $errorMsg.html result.error_messages.join '<br>'
  else
    $errorMsg.text $(formErrorSelector).attr 'data-default-error-message'
  do $(formErrorSelector).show

toggleSubmitting = (enable) ->
  $(contactFormFieldsetSelector).prop 'disabled', enable
  $(contactFormLoadingImgSelector).toggle enable

$ ->
  do setupToggleSidebarCollapse
  do setupContactModal
