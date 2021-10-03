const {
  OctyEventManager,
  OctyContentManager,
  OctyRecommendationsManager,

  recommendations,
  functions,
  content,
  events,

  endpoints,
  DSCMap,
  utils
} = require('./src')

async function main () {
  try {
    // Init customer session
    const octyCustomerID = functions.getLSValues()
    const customerId = await functions.createUpdateProfile(octyCustomerID)
    functions.setLSValues(customerId)
  } catch (error) {
    console.error(error)
    functions.clearLSValues()
  }
  try {
    // events
    await OctyEventManager.registerEventListeners()
    // dynamic content
    await OctyContentManager.assessDynamicContentSections()
    // product recommendations
    await OctyRecommendationsManager.loadProductRecommendations()
  } catch (error) {
    console.error(error)
  }
}

jQuery(document).ready(function () {
  console.log('loaded')
  main()
})

console.log('loaded')
