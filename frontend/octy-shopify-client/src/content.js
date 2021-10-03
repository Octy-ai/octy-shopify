const endpoints = require('./config/endpoints')
const DCSMap = require('./config/DCSMap')

function getDynamicContent (octyCustomerID) {
  return new Promise((resolve) => {
    const xhr = new XMLHttpRequest()
    xhr.open('POST', endpoints.getContentURI, true)
    xhr.onload = function (e) {
      if (xhr.readyState === 4) {
        if (xhr.status === 200) {
          resolve([true, JSON.parse(xhr.responseText).sections])
        } else {
          console.warn('Failed to get dynamic content!')
          resolve([false, null])
        }
      }
    }
    xhr.onerror = function (e) {
      console.warn('Failed to get dynamic content!')
      resolve([false, null])
    }
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(
      JSON.stringify({
        octy_customer_id: octyCustomerID,
        sections: DCSMap
      })
    )
  })
}

function loadDefaultContent () {
  const sections = document.querySelectorAll('[data-octy-content]')
  if (sections != null) {
    for (const i in sections) {
      if (sections.hasOwnProperty(i)) {
        const section = sections[i]
        const sectionID = section.getAttribute('data-octy-content')
        while (section.firstChild) {
          section.removeChild(section.firstChild)
        }
        let defaultValue
        for (const j in DCSMap) {
          if (DCSMap[j].section_id === sectionID) {
            defaultValue = DCSMap[j].default_value
            break
          }
        }
        section.innerHTML += defaultValue
      }
    }
    console.log('Default content loaded!')
  } else {
    console.log('No dynamic sections found!')
  }
}

function loadDynamicContent (dynamicSections) {
  const sections = document.querySelectorAll('[data-octy-content]')
  if (sections != null) {
    for (const i in sections) {
      if (sections.hasOwnProperty(i)) {
        const section = sections[i]
        const sectionID = section.getAttribute('data-octy-content')
        while (section.firstChild) {
          section.removeChild(section.firstChild)
        }
        let content
        for (const j in dynamicSections) {
          if (dynamicSections[j].sectionID === sectionID) {
            content = dynamicSections[j].content
            break
          }
        }
        section.innerHTML += content
      }
    }
  }
}

module.exports = {
  getDynamicContent,
  loadDefaultContent,
  loadDynamicContent
}
