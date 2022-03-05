import { render, screen } from '@testing-library/react';
import Home from '../components/Home';

describe('Header component', () => {
 test('it renders', () => {
   render(<Home />);
   expect(screen.getByText('SHOWBASE')).toBeInTheDocument();
   expect(screen.getByText('PROFILE')).toBeInTheDocument();
   expect(screen.getByTestId('HOME') ).toBeInTheDocument();
   expect(screen.getByTestId('ABOUT') ).toBeInTheDocument();
   expect(screen.getByTestId('MOVIES & SHOWS') ).toBeInTheDocument();
//    expect(screen.getByTestId('summary') ).toBeInTheDocument();
//    expect(screen.getByTestId('posttype') ).toBeInTheDocument();
 });
})