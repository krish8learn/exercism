def leap_year(year):
    """
    Determine if a year is a leap year.
    
    A leap year is:
    - Divisible by 400, OR
    - Divisible by 4 AND not divisible by 100
    """
    return year % 400 == 0 or (year % 4 == 0 and year % 100 != 0)
