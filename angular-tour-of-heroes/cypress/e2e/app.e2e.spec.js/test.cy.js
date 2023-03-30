describe('Tour of Heroes App Testing', () => {
    it(`Should have the title 'Fortune Cookie'`, () => {
        cy.visit('localhost:4200')
        cy.contains('Fortune Cookie');
    });
});