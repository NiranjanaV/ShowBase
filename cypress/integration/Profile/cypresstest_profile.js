///<reference types="cypress"/>
 
it('Profile Test', function(){
    cy.viewport(1525, 836)
    cy.visit('http://localhost:3000/UserProfile')
    cy.get(':nth-child(3) > a').click()
    cy.get('input').type('Star Wars{enter}')
    cy.get(':nth-child(2) > a').click()
    cy.get('ul > :nth-child(1) > a').click()
    cy.get('a > h1').click()
})