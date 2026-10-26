import Heap from "./heap";

function main() {
  const h = new Heap();

  const values = [5, 3, 8, 1, 9, 2, 7, 4, 6, 0];
  console.log("Inserting:", values);
  for (const v of values) {
    h.insert(v);
  }

  console.log("Internal array:", [...h.data]);

  console.log("\nExtracting min one by one:");
  const extracted: number[] = [];
  while (h.data.length > 0) {
    extracted.push(h.data[0]); // root is the current min
    h.delete();
  }
  console.log("Extracted order:", extracted);

  const sorted = [...values].sort((a, b) => a - b);
  console.log(
    "\nIs it sorted ascending? ",
    JSON.stringify(extracted) === JSON.stringify(sorted)
  );

  console.log("\nEdge cases:");
  console.log("  delete on empty heap ->", h.delete());

  const single = new Heap();
  single.insert(42);
  single.delete();
  console.log("  heap empty after single delete ->", single.data.length === 0);
}

main();
