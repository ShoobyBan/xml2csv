# xml2csv
XML to CSV converter, supports missing tags.
Data format is flexible, currently:
  
    <catalog>
      <product>
        <*any1*>A</*any1*>
        <*any2*>B</*any2*>
      </product>
      <product>
        <*any2*>C</*any2*>
        <*any3*>D</*any3*>
      </product>
    </catalog>

Meaning inside product on one level the content can be anything.

CSV will keep the order of the tags.

    any1,any2,any3
    A   ,B   ,
        ,C   ,D

Output is written to standard output, input is the first argument.
