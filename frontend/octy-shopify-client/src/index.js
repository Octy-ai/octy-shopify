const endpoints = require('./config/endpoints')
const DCSMap = require('./config/DCSMap')
const utils = require('./utils')

const recommendations = require('./recommendations')
const functions = require('./functions')
const content = require('./content')
const events = require('./events')

// This flag is used to determine if any requests to octy-shopify Shopify app should be made.
// If false, octy-customer-id has not been set in Local Storage.
window.isProfileEnabled = false

class OctyEventManager {
  static async registerEventListeners () {
    // Add a 'data-octy-event-{type}="{eventType}"' attribute on all HTML elements that should trigger an event instance.

    // register click event listeners on HTML elements with : data-octy-event-click="{eventType}"
    // example :: data-octy-event-click="showed_love"
    const clickEls = document.querySelectorAll('[data-octy-event-click]')
    if (clickEls != null) {
      for (const i in clickEls) {
        if (clickEls.hasOwnProperty(i)) {
          const cEventType = clickEls[i].getAttribute('data-octy-event-click')
          clickEls[i].addEventListener(
            'click',
            function () {
              OctyEventManager.capture(cEventType)
            },
            false
          )
        }
      }
      console.log('Click octy-event handlers registred!')
    } else {
      console.warn('No click octy-event handlers registred!')
    }

    // register hover event listeners on HTML elements with : data-octy-event-hover="{eventType}"
    // example :: data-octy-event-hover="engaged_with_content"
    const hoverEls = document.querySelectorAll('[data-octy-event-hover]')
    if (hoverEls != null) {
      for (const i in hoverEls) {
        if (hoverEls.hasOwnProperty(i)) {
          const hEventType = hoverEls[i].getAttribute('data-octy-event-hover')
          hoverEls[i].addEventListener(
            'mouseenter',
            function () {
              OctyEventManager.capture(hEventType)
            },
            false
          )
        }
      }
      console.log('Hover octy-event handlers registred!')
    } else {
      console.warn('No hover octy-event handlers registred!')
    }
    // NOTE: Most JS event listeners can be registered to capture
    // octy event instances : https://developer.mozilla.org/en-US/docs/Web/Events
  }

  static async capture (eventType) {
    if (window.isProfileEnabled) {
      console.log(
        'Event: ' + eventType + ' occurred. Processing event capture...'
      )
      let validEvent = false
      // switch through event type to specify event properties
      // that are associated with this event instance
      let eventProperties
      switch (eventType) {
        // Specify event properties for each event type
        case 'showed_love':
          eventProperties = {
            device: utils.getUA()
            // ... add other properties here
          }
          validEvent = true
          break
        case 'engaged_with_content':
          eventProperties = {
            medium: 'web content',
            platform: utils.getUA()
            // ... add other properties here
          }
          validEvent = true
          break
        default:
          break
      }
      if (validEvent) {
        const octyCustomerID = functions.getLSValues()
        await events.createOctyEvent(octyCustomerID, eventType, eventProperties)
      }
    } else {
      console.log(
        'Event: ' +
             eventType +
             ' occurred but no credentials set in LS. Unable to capture this event instance.'
      )
    }
  }
}

class OctyContentManager {
  static async assessDynamicContentSections () {
    const sections = document.querySelectorAll('[data-octy-content]')
    if (window.isProfileEnabled) {
      const octyCustomerID = await functions.getLSValues()
      const dynamicSections = await content.getDynamicContent(octyCustomerID)
      if (dynamicSections[0]) {
        // dynamicSections[0] is bool result
        content.loadDynamicContent(dynamicSections[1])
      } else {
        // some error has occurred
        content.loadDefaultContent(sections)
      }
    } else {
      // load default content
      content.loadDefaultContent(sections)
    }
  }
}

class OctyRecommendationsManager {
  static async loadProductRecommendations () {
    const row = document.getElementById('octy-rec-products')
    const recHeader = document.getElementById('rec-header')
    row.style.display = 'none'
    recHeader.style.display = 'none'
    const recLimit = 3 // the number of product recommendations you wish to render.
    // Note, backend octy-shopify app only returns 10 per request.
    if (window.isProfileEnabled) {
      const octyCustomerID = await functions.getLSValues()
      const data = await recommendations.getProductRecommendations(octyCustomerID)
      if (data.length < 1) {
        // If no items are returned, hide the product recommendation section.
        row.style.display = 'none'
        recHeader.style.display = 'none'
      } else {
        recHeader.style.display = 'block'
        row.style.display = 'block'
        let recCount = 0
        for (const item of data) {
          if (recCount >= recLimit) {
            break
          }
          row.insertAdjacentHTML(
            'afterbegin',
            ' <div class="column"> <div class="card"> <img src="' +
                 item.item_image_url +
                 '" style="width:100%"> <h3>' +
                 item.item_name +
                 '</h3> <p class="price">£' +
                 item.item_price +
                 '</p> <p><form><button formaction="' +
                 item.item_link +
                 '">View Product</button></form></p> </div> </div>'
          )
          recCount++
        }
      }
    } else {
      // if profile is not enabled, hide the product recommendation section.
      row.style.display = 'none'
      recHeader.style.display = 'none'
    }
  }
}

module.exports = {
  OctyEventManager,
  OctyContentManager,
  OctyRecommendationsManager,

  recommendations,
  functions,
  content,
  events,

  endpoints,
  DCSMap,
  utils
}
