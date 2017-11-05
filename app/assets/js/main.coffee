$ = window.$ = window.jQuery = require 'jquery'
require 'bootstrap'

collapseSelector = '.js-vocanew-collapse'

setupToggleSidebarCollapse = ->
  $(collapseSelector).on 'click touchstart', '.js-vocanew-collapse-btn', (ev) ->
    do ev.preventDefault
    $(collapseSelector).toggleClass 'vocanew-collapse-active'

$ ->
  do setupToggleSidebarCollapse
