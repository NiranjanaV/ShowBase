///<reference types="cypress"/>
 //test for movies and shows page
 // All functionalities
 
it('Movies and Shows test', function(){
    cy.viewport(1525, 836)
    cy.visit('http://localhost:3000/')
    cy.get('ul > :nth-child(2) > a').click()

    cy.get(':nth-child(3) > input').type('niranjana')
    cy.get(':nth-child(4) > input').type('niranjana')
    cy.get(':nth-child(5) > input').dblclick()
    cy.get('ul > :nth-child(2) > a').click()
    
    cy.get(':nth-child(1) > .list > :nth-child(1) > .reusable').click()

    cy.get(':nth-child(6) > form > .input-container > button').click()
    cy.get('a > h1').click()
    cy.get('ul > :nth-child(2) > a').click()
    cy.get('li > a').click()
    
    cy.get('input').type('harry')
    cy.get('a > h1').click()

})