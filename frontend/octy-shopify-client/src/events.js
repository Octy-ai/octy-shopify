const endpoints = require('./config/endpoints')

function createOctyEvent (octyCustomerID, eventType, eventProperties) {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()
    xhr.open('POST', endpoints.createEventURI, true)
    xhr.onload = function (e) {
      if (xhr.readyState === 4) {
        if (xhr.status === 201) {
          console.log('Captured Octy event instance! :' + eventType)

          resolve('done')
        } else {
          console.warn('Failed to capture Octy event instance!')

          reject(new Error('Server Error'))
        }
      }
    }
    xhr.onerror = function (e) {
      console.warn('Failed to capture Octy event instance!')

      reject(new Error('Server Error'))
    }

    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(
      JSON.stringify({
        octy_customer_id: octyCustomerID,
        event_type: eventType,
        event_properties: eventProperties
      })
    )
  })
}

module.exports = {
  createOctyEvent
}
