def to_rna(dna_strand):
    rna = ""
    translate = {"G" : "C", "C" : "G", "T" : "A", "A" : "U"}
    for ch in dna_strand:
        if ch in translate:
            rna += translate[ch]
    return rna
