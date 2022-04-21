///<reference types="cypress"/>
 //test for Account page, i.e my profile

 // All functionalities
 
it('Profile Test', function(){

    cy.viewport(1525, 836)
    cy.visit('http://localhost:3000/login')
    cy.get(':nth-child(3) > input').type('niranjana')

    cy.get(':nth-child(4) > input').type('niranjana')
    cy.get(':nth-child(5) > input').dblclick()
    cy.get(':nth-child(3) > a').click()
    cy.get('[href="/UserProfile/watched"] > .service-item').click()

    cy.go('back')
    cy.get('[href="/UserProfile/Watchlist"] > .service-item').click()
    cy.go('back')
    cy.get('ul > :nth-child(1) > a').click()
    cy.get('ul > :nth-child(1) > a').click()

    cy.get(':nth-child(1) > .resultscard > .movieinfo > .add > :nth-child(1)').click()
    cy.go('back')
    cy.get(':nth-child(1) > .resultscard > .movieinfo > .add > :nth-child(2)').click()
    cy.go('back')
    cy.go('back')

    cy.get(':nth-child(2) > a').click()
    cy.get('input').type('s')
    cy.get(':nth-child(1) > .resultscard > .movieinfo > .add > .button').click()





})