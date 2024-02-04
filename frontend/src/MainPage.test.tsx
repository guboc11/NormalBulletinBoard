import { render, screen } from '@testing-library/react';
import MainPage from './MainPage';

test("메인페이지가 존재함",()=>{
  render(<MainPage/>)
  const mainpage = screen.getByTitle("MainPage")
  expect(mainpage).toBeInTheDocument();
})