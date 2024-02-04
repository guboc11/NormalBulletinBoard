import { render, screen } from '@testing-library/react';
import MainBody from './MainBody';

test("MainBody가 존재한다.", ()=>{
  render(<MainBody/>)
  const mainbody = screen.getByTitle("MainBody")
  expect(mainbody).toBeInTheDocument();
})