const endpoints = require('./config/endpoints')

async function getProductRecommendations (octyCustomerID) {
  return new Promise((resolve, reject) => {
    const row = document.getElementById('octy-rec-products')
    const recHeader = document.getElementById('rec-header')
    const xhr = new XMLHttpRequest()
    xhr.open('POST', endpoints.getRecommendationsURI, true)
    xhr.onload = function (e) {
      if (xhr.readyState === 4) {
        if (xhr.status === 200) {
          resolve(JSON.parse(xhr.responseText).items)
        } else {
          console.warn('Failed to get product recommendations!')
          row.style.display = 'none'
          recHeader.style.display = 'none'

          reject(new Error('Server Error'))
        }
      }
    }
    xhr.onerror = function (e) {
      console.warn('Failed to get product recommendations!')
      row.style.display = 'none'
      recHeader.style.display = 'none'

      reject(new Error('Server Error'))
    }
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(
      JSON.stringify({
        octy_customer_id: octyCustomerID
      })
    )
  })
}

module.exports = {
  getProductRecommendations
}
