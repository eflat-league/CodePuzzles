import unittest
from first_task import rotate

class TestRotate(unittest.TestCase):
    def test_zero_rotation(self):
        self.assertEqual(rotate(50, 'L', 0), 50)
        self.assertEqual(rotate(50, 'R', 0), 50)
        self.assertEqual(rotate(0, 'L', 0), 0)
        self.assertEqual(rotate(99, 'R', 0), 99)

    def test_left_rotation(self):
        self.assertEqual(rotate(50, 'L', 1), 49)
        self.assertEqual(rotate(50, 'L', 50), 0)
        self.assertEqual(rotate(50, 'L', 51), 99)
        self.assertEqual(rotate(50, 'L', 75), 75)
        self.assertEqual(rotate(50, 'L', 100), 50)
        self.assertEqual(rotate(50, 'L', 321), 29)

    def test_right_rotation(self):
        self.assertEqual(rotate(50, 'R', 1), 51)
        self.assertEqual(rotate(50, 'R', 50), 0)
        self.assertEqual(rotate(50, 'R', 51), 1)
        self.assertEqual(rotate(50, 'R', 75), 25)
        self.assertEqual(rotate(50, 'R', 100), 50)
        self.assertEqual(rotate(50, 'R', 321), 71)


if __name__ == '__main__':
    unittest.main()