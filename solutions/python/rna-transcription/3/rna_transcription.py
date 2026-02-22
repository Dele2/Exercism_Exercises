def to_rna(dna_strand):
    rna = ""
    translate_to_rna = {"G" : "C", "C" : "G", "T" : "A", "A" : "U"}
    for ch in dna_strand:
        if ch in translate_to_rna:
            rna += translate_to_rna[ch]
    return rna
