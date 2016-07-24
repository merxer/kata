jest.unmock('../sum');

describe('sum', () => {
  const sum = require('../sum.js');
  it('adds 1 + 2 to equal 3', () => {
    expect(sum(1,2)).toBe(3);
  });

  it('input is string', () => {
    expect(sum('a','b')).toBe('ab');
  });
});
