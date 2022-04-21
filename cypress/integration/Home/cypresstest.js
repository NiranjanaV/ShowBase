///<reference types="cypress"/>
 //test for home page
 // All functionalities
 
it('Home test', function(){
    cy.viewport(1525, 836)
    cy.visit('http://localhost:3000/')
    cy.get('a > h1').click()

    cy.get('a > img').click()
    cy.go('back')

    cy.get('ul > :nth-child(1) > a').click()

    cy.get('ul > :nth-child(2) > a').click()
    cy.go('back')

    cy.get(':nth-child(3) > a').click()
    cy.go('back')

    
})