import { render, screen } from '@testing-library/react';
import Home from '../components/Home';
import {
  BrowserRouter as Router,
  Route,
  Routes
} from 'react-router-dom';


describe('Header component', () => {
 test('it renders', () => {
   render(
    <Router>
    <Routes>
      <Route path="/" element={<Home />} />
    </Routes>
  </Router>
   );
  //  expect(screen.getByText('Contact')).toBeInTheDocument();
 });
})