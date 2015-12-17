# xml2csv
XML to CSV converter, supports missing tags
Data format is flexible, currently it's 
  
    <catalog>
      <product>
        <*any1*>*anything*</*any1*>
        <*any2*>*anything*</*any2*>
      </product>
      <product>
        <*any2*>*any*</*any2*>
        <*any3*>*any*</*any3*>
      </product>
    </catalog>
  
Output is written to standard output, input is the first argument.
