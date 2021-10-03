const endpoints = require('./config/endpoints')
const utils = require('./utils')

function getLSValues () {
  const octyCustomerID = localStorage.getItem('octy-customer-id')
  return octyCustomerID
}

function setLSValues (octyCustomerID) {
  // set octy-customer-id in local storage
  localStorage.setItem('octy-customer-id', octyCustomerID)
}

function clearLSValues () {
  localStorage.removeItem('octy-customer-id')
}

function createUpdateProfile (octyCustomerID) {
  return new Promise((resolve, reject) => {
    let octyCustomerId // customer ID sent in request body.
    let generatedNewId = false // bool flag to indicate to server
    // if octy-customer-id was found in local storage.
    const profileData = {} // profile data sent in request body.
    let shopifyCustomerID = window.ShopifyAnalytics.meta.page.customerId

    if (octyCustomerID == null) {
      // Check octy-customer-id in Local Storage.
      // octy-customer-id not in Local Storage
      // Generate new octy-customer-id
      octyCustomerId = 'octy-customer-id-' + utils.uuidv4()
      generatedNewId = true
    } else {
      // octy-customer-id found in Local Storage
      // Get octy-customer-id from Local Storage
      octyCustomerId = octyCustomerID
    }

    if (shopifyCustomerID == null) {
      // Check if customer is authenticated
      // Customer is not authenticated
      shopifyCustomerID = '' // convert to empty string for request body
    } else {
      // Customer is authenticated
      profileData.shopify_customer_id = shopifyCustomerID
    }

    // < HTTP request >

    const xhr = new XMLHttpRequest()
    xhr.open('POST', endpoints.createCustomerURI, true)
    xhr.onload = function (e) {
      if (xhr.readyState === 4) {
        if (xhr.status === 201) {
          window.isProfileEnabled = true
          console.log('Created/Updated Octy profile!')

          resolve(JSON.parse(xhr.responseText).customer_id)
        } else {
          window.isProfileEnabled = false
          console.warn(
            'Failed to create new Octy profile! Unable to track events or provide personalised experience during this session'
          )

          reject(new Error('Server Error'))
        }
      }
    }
    xhr.onerror = function (e) {
      window.isProfileEnabled = false
      console.warn(
        'Failed to create new Octy profile! Unable to track events or provide personalised experience during this session'
      )

      reject(new Error('Server Error'))
    }
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(
      JSON.stringify({
        octy_customer_id: octyCustomerId,
        shopify_customer_id: shopifyCustomerID,
        has_charged: false,
        profile_data: profileData,
        platform_info: {
          user_agent: utils.getUA()
        }, // TODO: Get some other platform info
        generated_new_id: generatedNewId
      })
    )
  })
}


module.exports = {
  getLSValues,
  setLSValues,
  clearLSValues,
  createUpdateProfile
}
